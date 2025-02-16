// design/design.go
package design

import (
	"aion/design/services"

	. "goa.design/goa/v3/dsl"
)

var _ = API("ausa_service", func() {
	Title("Ausa Service")
	Description("A service for something that need to be managed")
	Server("ausa_service", func() {
		Host("localhost", func() {
			URI("http://localhost:9090")
		})
	})

	HTTP(func() {
		Path("/api/v1")
	})

})

// commmon
// - Payload
//var paginationPayload = payloads.PaginationPayload

// - Response
//var paginationResponse = responses.PaginatedResponse
//var successResponse = responses.SuccessResponse
//var errorResponse = responses.ErrorResponse

// User
// - Service
var userService = services.UserService

// Location
//var locationService = services.LocationService

// Store
// - FileStore
var fileStore = services.FileStoreService
