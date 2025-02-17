package main

import (
	fileStore "aion/gen/file_store"
	fileStoreSvr "aion/gen/http/file_store/server"
	userGenSvr "aion/gen/http/user/server"
	userGen "aion/gen/user"
	"aion/internal/config"
	"aion/internal/utils"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"

	"goa.design/clue/debug"
)

// HandleHttpServer initializes and starts the HTTP server, handling GraphQL, REST, and Swagger endpoints.
// It also sets up CORS, error handling, and optional debug endpoints.
func HandleHttpServer(ctx context.Context, u *url.URL, wg *sync.WaitGroup, errc chan error, dbg bool, epsMap map[config.EndpointName]interface{}) {
	var (
		dec     = goahttp.RequestDecoder  // Decoder for HTTP requests
		enc     = goahttp.ResponseEncoder // Encoder for HTTP responses
		mux     goahttp.Muxer             // HTTP request multiplexer (router)
		handler http.Handler              // Final HTTP handler
	)
	{
		mux = goahttp.NewMuxer() // Initialize the request multiplexer

		// Enable debug endpoints if debug mode is on
		if dbg {
			debug.MountPprofHandlers(debug.Adapt(mux))   // Profiling endpoints for performance analysis
			debug.MountDebugLogEnabler(debug.Adapt(mux)) // Enable/disable debug logs dynamically
		}

		// Set up error handler
		eh := errorHandler(ctx)

		// Mount RESTful service endpoints
		mountServices(ctx, mux, dec, enc, eh, epsMap)

		// Health check endpoint for monitoring server health
		mux.Handle("GET", "/healthz", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				_, _ = w.Write([]byte("Method Not Allowed"))
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		})

		handler = mux // Assign the configured mux to handler
	}

	// Wrap handler with additional middlewares: error handling, CORS, logging, and debugging
	handler = withErrorHandler(handler, ctx) // Error handling middleware
	handleMuxer(mux)                         // Swagger and ReDoc setup
	handler = enableCORS(handler)            // Enable Cross-Origin Resource Sharing
	if dbg {
		handler = debug.HTTP()(handler) // Add debug HTTP middleware if in debug mode
	}
	handler = log.HTTP(ctx)(handler) // Add logging middleware for request logs

	// Create and configure HTTP server with timeouts and handler
	srv := &http.Server{
		Addr:              u.Host,
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 60, // Timeout for reading HTTP headers
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a goroutine
		go func() {
			log.Printf(ctx, "Starting server on %s", u.Host)
			errc <- srv.ListenAndServe()
		}()

		// Graceful shutdown on context cancellation
		<-ctx.Done()
		log.Printf(ctx, "Shutting down server at %q", u.Host)

		// Context with timeout for shutdown process
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

// withErrorHandler wraps a handler to recover from panics and log errors
func withErrorHandler(handler http.Handler, logCtx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}
				utils.Log.Error(logCtx, w, err)
				log.Printf(logCtx, "ERROR: %s", err.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}

// enableCORS enables Cross-Origin Resource Sharing for all endpoints,
// allowing specified methods and headers.
func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// errorHandler handles errors for HTTP requests, logging errors with contextual details.
func errorHandler(logCtx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		utils.Log.Error(logCtx, w, err)
		log.Printf(logCtx, "ERROR: %s", err.Error())
	}
}

// mountServices mounts RESTful service endpoints, mapping them from the endpoint map.
func mountServices(ctx context.Context, mux goahttp.Muxer, dec func(*http.Request) goahttp.Decoder, enc func(context.Context, http.ResponseWriter) goahttp.Encoder, eh func(context.Context, http.ResponseWriter, error), epsMap map[config.EndpointName]interface{}) {
	var fileStoreServer *fileStoreSvr.Server
	var userGenServer *userGenSvr.Server
	// Mount each endpoint based on the provided endpoint map
	for name, eps := range epsMap {
		switch name {
		case config.StoreEndPoint:
			FileStoreEndpoints := eps.(*fileStore.Endpoints)
			fileStoreServer = fileStoreSvr.New(FileStoreEndpoints, mux, dec, enc, eh, nil)
			fileStoreSvr.Mount(mux, fileStoreServer)
		case config.UserEndPoint:
			userEndpoints := eps.(*userGen.Endpoints)
			userGenServer = userGenSvr.New(userEndpoints, mux, dec, enc, eh, nil)
			userGenSvr.Mount(mux, userGenServer)
		}

	}
}

// handleMuxer sets up additional handlers for Swagger and ReDoc documentation.
func handleMuxer(mux goahttp.Muxer) {
	// Serve Swagger UI static files
	fs := http.FileServer(http.Dir("./swagger-ui"))
	mux.Handle("GET", "/swagger-ui/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger-ui/", fs).ServeHTTP(w, r)
	})

	// Serve OpenAPI specification file
	mux.Handle("GET", "/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/openapi3.yaml")
	})

	// Serve Swagger HTML documentation
	mux.Handle("GET", "/docs/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		swagger, err := swagger()
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(swagger)
	})

	// Serve ReDoc HTML documentation
	mux.Handle("GET", "/docs/redoc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		redoc, err := redoc()
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(redoc)
	})
}

// redoc loads the ReDoc HTML file from the static directory for documentation.
func redoc() ([]byte, error) {
	filePath := "./static/redoc.html"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error Reading Redoc: %v\n", err)
		return nil, err
	}
	return fileContent, nil
}

// swagger loads the Swagger HTML file from the static directory for API documentation.
func swagger() ([]byte, error) {
	filePath := "./static/swagger.html"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error Reading Redoc: %v\n", err)
		return nil, err
	}
	return fileContent, nil
}
