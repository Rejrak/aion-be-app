// design/design.go
package design

import (
	"aion/design/services"

	. "goa.design/goa/v3/dsl"
)

var _ = API("aion_service", func() {
	Title("Aion Service")
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

// User
// - Service
var _ = services.UserService
