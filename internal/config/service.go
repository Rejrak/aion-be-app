package config

import (
	userGen "aion/gen/user"
	"aion/internal/database/db"
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
	userConfig := initializedUserService()
	epsMap := make(map[EndpointName]interface{})
	db.ConnectDb()

	services := []ServiceConfig{userConfig}
	for _, serviceConfig := range services {
		svc := serviceConfig.NewService()              // Create a new service instance
		endpoints := serviceConfig.NewEndpoints(svc)   // Generate endpoints for the service
		epsMap[serviceConfig.EndpointName] = endpoints // Add the endpoints to the map with the endpoint name as the key
	}

	return epsMap // Return the map containing all initialized service endpoints
}
