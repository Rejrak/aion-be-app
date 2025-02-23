package main

import (
	"be/internal/config"
	"be/internal/utils"
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

func HandleHttpServer(ctx context.Context, u *url.URL, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var handler http.Handler
	var mux goahttp.Muxer = config.InitializeMuxer(ctx, dbg)
	mux = withDocsHandler(mux)
	handler = mux
	handler = withErrorHandler(handler, ctx)
	handler = withCORS(handler)

	if dbg {
		handler = debug.HTTP()(handler)
	}

	handler = log.HTTP(ctx)(handler)

	srv := &http.Server{
		Addr:              u.Host,
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 60,
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		go func() {
			log.Printf(ctx, "Starting server on %s", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		log.Printf(ctx, "Shutting down server at %q", u.Host)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

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

func withCORS(h http.Handler) http.Handler {
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

// withDocsHandler sets up HTTP handlers for serving Swagger UI, OpenAPI spec, and API documentation.
// It registers routes for serving static files and dynamically generated documentation pages.
func withDocsHandler(mux goahttp.Muxer) goahttp.Muxer {
	fs := http.FileServer(http.Dir("./swagger-ui"))
	mux.Handle("GET", "/swagger-ui/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger-ui/", fs).ServeHTTP(w, r)
	})

	mux.Handle("GET", "/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/openapi3.yaml")
	})

	mux.Handle("GET", "/docs/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		swagger, err := withSwagger()
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(swagger)
	})

	mux.Handle("GET", "/docs/redoc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		redoc, err := withRedoc()
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(redoc)
	})
	return mux
}

func withRedoc() ([]byte, error) {
	filePath := "./static/redoc.html"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error Reading Redoc: %v\n", err)
		return nil, err
	}
	return fileContent, nil
}

func withSwagger() ([]byte, error) {
	filePath := "./static/swagger.html"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error Reading Redoc: %v\n", err)
		return nil, err
	}
	return fileContent, nil
}
