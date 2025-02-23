package config

import (
	userGenSvr "be/gen/http/user/server"
	userGen "be/gen/user"
	userService "be/internal/user"
	"be/internal/utils"
	"context"
	"net/http"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

type EndpointName string

const (
	StoreEndPoint EndpointName = "store"
	UserEndPoint  EndpointName = "user"
)

type ServiceConfig struct {
	EndpointName EndpointName
	NewService   func() interface{}
	NewEndpoints func(svc interface{}) interface{}
	Mount        func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error))
}

func withUserService() ServiceConfig {
	return ServiceConfig{
		EndpointName: UserEndPoint,
		NewService:   func() interface{} { return userService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			userSvc := svc.(userGen.Service)
			endpoints := userGen.NewEndpoints(userSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			userEndpoints := endpoints.(*userGen.Endpoints)
			server := userGenSvr.New(userEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			userGenSvr.Mount(mux, server)
		},
	}
}

func InitializeMuxer(ctx context.Context, enableDebug bool) goahttp.Muxer {
	mux := goahttp.NewMuxer()
	eh := errorHandler(ctx)

	mountServices(mux, eh)

	if enableDebug {
		debug.MountPprofHandlers(debug.Adapt(mux))
		debug.MountDebugLogEnabler(debug.Adapt(mux))
	}

	mux.Handle("GET", "/healthz", healthCheckHandler)

	return mux
}

func mountServices(mux goahttp.Muxer, eh func(context.Context, http.ResponseWriter, error)) {
	services := []ServiceConfig{
		withUserService(),
	}

	for _, service := range services {
		svc := service.NewService()
		endpoints := service.NewEndpoints(svc)
		service.Mount(mux, endpoints, eh)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method Not Allowed"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func errorHandler(ctx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(_ context.Context, w http.ResponseWriter, err error) {
		utils.Log.Error(ctx, w, err)
		log.Printf(ctx, "ERROR: %s", err.Error())
	}
}
