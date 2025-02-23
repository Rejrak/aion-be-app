package config

import (
	trainingPlanGenSvr "be/gen/http/trainingplan/server"
	trainingPlanGen "be/gen/trainingplan"
	trainingPlanService "be/internal/features/trainingplan"

	userGenSvr "be/gen/http/user/server"
	userGen "be/gen/user"
	userService "be/internal/features/user"

	workoutGenSvr "be/gen/http/workout/server"
	workoutGen "be/gen/workout"
	workoutService "be/internal/features/workout"

	workoutExerciseGenSvr "be/gen/http/workoutexercise/server"
	workoutExerciseGen "be/gen/workoutexercise"
	workoutExerciseService "be/internal/features/workoutexercise"

	muscleGroupGenSvr "be/gen/http/musclegroup/server"
	muscleGroupGen "be/gen/musclegroup"
	muscleGroupService "be/internal/features/musclegroup"

	workoutProgressGenSvr "be/gen/http/workoutprogress/server"
	workoutProgressGen "be/gen/workoutprogress"
	workoutProgressService "be/internal/features/workoutprogress"

	exerciseProgressGen "be/gen/exerciseprogress"
	exerciseProgressGenSvr "be/gen/http/exerciseprogress/server"
	exerciseProgressService "be/internal/features/exerciseprogress"

	"be/internal/utils"
	"context"
	"net/http"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

type EndpointName string

const (
	userEndPoint            EndpointName = "user"
	trainingPlanEndPoint    EndpointName = "trainingplan"
	workoutEndPoint         EndpointName = "workout"
	muscleGroupEndPoint     EndpointName = "musclegroup"
	workoutExerciseEndPoint EndpointName = "workoutexercise"
	workoutProgressEndPoint EndpointName = "workoutprogress"
	exerciseProgress        EndpointName = "exerciseprogress"
)

type ServiceConfig struct {
	EndpointName EndpointName
	NewService   func() interface{}
	NewEndpoints func(svc interface{}) interface{}
	Mount        func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error))
}

func withUserService() ServiceConfig {
	return ServiceConfig{
		EndpointName: userEndPoint,
		NewService:   func() interface{} { return userService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(userGen.Service)
			endpoints := userGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*userGen.Endpoints)
			server := userGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			userGenSvr.Mount(mux, server)
		},
	}
}

func withTrainingPlanService() ServiceConfig {
	return ServiceConfig{
		EndpointName: trainingPlanEndPoint,
		NewService:   func() interface{} { return trainingPlanService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(trainingPlanGen.Service)
			endpoints := trainingPlanGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*trainingPlanGen.Endpoints)
			server := trainingPlanGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			trainingPlanGenSvr.Mount(mux, server)
		},
	}
}

func withWorkoutService() ServiceConfig {
	return ServiceConfig{
		EndpointName: workoutEndPoint,
		NewService:   func() interface{} { return workoutService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(workoutGen.Service)
			endpoints := workoutGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*workoutGen.Endpoints)
			server := workoutGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			workoutGenSvr.Mount(mux, server)
		},
	}
}

func withMuscleGroupService() ServiceConfig {
	return ServiceConfig{
		EndpointName: muscleGroupEndPoint,
		NewService:   func() interface{} { return muscleGroupService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(muscleGroupGen.Service)
			endpoints := muscleGroupGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*muscleGroupGen.Endpoints)
			server := muscleGroupGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			muscleGroupGenSvr.Mount(mux, server)
		},
	}
}

func withWorkoutExerciseService() ServiceConfig {
	return ServiceConfig{
		EndpointName: workoutExerciseEndPoint,
		NewService:   func() interface{} { return workoutExerciseService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(workoutExerciseGen.Service)
			endpoints := workoutExerciseGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*workoutExerciseGen.Endpoints)
			server := workoutExerciseGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			workoutExerciseGenSvr.Mount(mux, server)
		},
	}
}

func withWorkoutProgressService() ServiceConfig {
	return ServiceConfig{
		EndpointName: workoutProgressEndPoint,
		NewService:   func() interface{} { return workoutProgressService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(workoutProgressGen.Service)
			endpoints := workoutProgressGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*workoutProgressGen.Endpoints)
			server := workoutProgressGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			workoutProgressGenSvr.Mount(mux, server)
		},
	}
}

func withExerciseProgressService() ServiceConfig {
	return ServiceConfig{
		EndpointName: exerciseProgress,
		NewService:   func() interface{} { return exerciseProgressService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			newSvc := svc.(exerciseProgressGen.Service)
			endpoints := exerciseProgressGen.NewEndpoints(newSvc)
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
		Mount: func(mux goahttp.Muxer, endpoints interface{}, eh func(context.Context, http.ResponseWriter, error)) {
			ep := endpoints.(*exerciseProgressGen.Endpoints)
			server := exerciseProgressGenSvr.New(ep, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, eh, nil)
			exerciseProgressGenSvr.Mount(mux, server)
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
		withTrainingPlanService(),
		withWorkoutService(),
		withMuscleGroupService(),
		withWorkoutExerciseService(),
		withWorkoutProgressService(),
		withExerciseProgressService(),
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
