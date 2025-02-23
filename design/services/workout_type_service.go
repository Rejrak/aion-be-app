package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var WorkoutType = Type("WorkoutType", func() {
	Attribute("id", String, "Unique ID of the workout type", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d482")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the workout type", func() {
		Example("Strength")
	})
	Attribute("description", String, "Description", func() {
		Example("Workout type focused on strength training")
	})
	Required("id", "name")
})

var CreateWorkoutTypePayload = Type("CreateWorkoutTypePayload", func() {
	Attribute("name", String, "Name of the workout type", func() {
		Example("Strength")
	})
	Attribute("description", String, "Description", func() {
		Example("Workout type focused on strength training")
	})
	Required("name")
})

var UpdateWorkoutTypePayload = Type("UpdateWorkoutTypePayload", func() {
	Attribute("id", String, "Workout Type ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d482")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the workout type", func() {
		Example("Strength Updated")
	})
	Attribute("description", String, "Description", func() {
		Example("Updated description")
	})
	Required("id", "name")
})

var WorkoutTypeService = Service("workouttype", func() {
	Description("Service for managing workout types")

	HTTP(func() {
		Path("/workouttypes")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Workout type not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new workout type")
		Payload(CreateWorkoutTypePayload)
		Result(WorkoutType)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a workout type by ID")
		Payload(func() {
			Attribute("id", String, "Workout Type ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d482")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(WorkoutType)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all workout types with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of workout types to return", func() {
				Example(10)
				Minimum(1)
				Maximum(100)
				Default(10)
			})
			Attribute("offset", Int, "Offset for pagination", func() {
				Example(0)
				Minimum(0)
				Default(0)
			})
		})
		Result(ArrayOf(WorkoutType))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a workout type")
		Payload(UpdateWorkoutTypePayload)
		Result(WorkoutType)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a workout type")
		Payload(func() {
			Attribute("id", String, "Workout Type ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d482")
				Format(FormatUUID)
			})
			Required("id")
		})
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})
