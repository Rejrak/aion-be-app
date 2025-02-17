package config

import (
	filestore "aion/gen/file_store"
	userGen "aion/gen/user"
	"aion/internal/database/db"
	storeService "aion/internal/store"
	userService "aion/internal/user"
	"context"

	"goa.design/clue/debug"
	"goa.design/clue/log"
)

type EndpointName string

const (
	StoreEndPoint EndpointName = "store"
	UserEndPoint  EndpointName = "user"
)

type ServiceConfig struct {
	EndpointName EndpointName                      // The name of the endpoint (used as a key in the map)
	NewService   func() interface{}                // Function to create a new service instance
	NewEndpoints func(svc interface{}) interface{} // Function to create endpoints for the service
}

func initializedStoreService() ServiceConfig {
	return ServiceConfig{
		EndpointName: StoreEndPoint,                                           // Label for this service
		NewService:   func() interface{} { return storeService.NewService() }, // Function to instantiate the store service
		NewEndpoints: func(svc interface{}) interface{} {
			// Generate new endpoints for the store service with debugging and logging middleware
			endpoints := filestore.NewEndpoints(svc.(filestore.Service))
			endpoints.Use(debug.LogPayloads()) // Log request and response payloads for debugging
			endpoints.Use(log.Endpoint)        // Log endpoint activities for monitoring
			return endpoints
		},
	}
}

func initializedUserService() ServiceConfig {
	return ServiceConfig{
		EndpointName: UserEndPoint,
		NewService:   func() interface{} { return userService.NewService() },
		NewEndpoints: func(svc interface{}) interface{} {
			endpoints := userGen.NewEndpoints(svc.(userGen.Service))
			endpoints.Use(debug.LogPayloads())
			endpoints.Use(log.Endpoint)
			return endpoints
		},
	}
}

func InitializeServices(ctx context.Context) map[EndpointName]interface{} {
	storeConfig := initializedStoreService()
	userConfig := initializedUserService()
	epsMap := make(map[EndpointName]interface{})
	db.ConnectDb()

	services := []ServiceConfig{storeConfig, userConfig}
	for _, serviceConfig := range services {
		svc := serviceConfig.NewService()              // Create a new service instance
		endpoints := serviceConfig.NewEndpoints(svc)   // Generate endpoints for the service
		epsMap[serviceConfig.EndpointName] = endpoints // Add the endpoints to the map with the endpoint name as the key
	}

	return epsMap // Return the map containing all initialized service endpoints
}
