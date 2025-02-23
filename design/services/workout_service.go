package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var Workout = Type("Workout", func() {
	Attribute("id", String, "Unique ID of the workout", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the workout", func() {
		Example("Workout A")
	})
	Attribute("trainingPlanId", String, "Training Plan ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Required("id", "name", "trainingPlanId")
})

var CreateWorkoutPayload = Type("CreateWorkoutPayload", func() {
	Attribute("name", String, "Name of the workout", func() {
		Example("Workout A")
	})
	Attribute("trainingPlanId", String, "Training Plan ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Required("name", "trainingPlanId")
})

var UpdateWorkoutPayload = Type("UpdateWorkoutPayload", func() {
	Attribute("id", String, "Workout ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the workout", func() {
		Example("Workout A Updated")
	})
	Attribute("trainingPlanId", String, "Training Plan ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Required("id", "name", "trainingPlanId")
})

var WorkoutService = Service("workout", func() {
	Description("Service for managing workouts")

	HTTP(func() {
		Path("/workouts")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Workout not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new workout")
		Payload(CreateWorkoutPayload)
		Result(Workout)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a workout by ID")
		Payload(func() {
			Attribute("id", String, "Workout ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(Workout)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all workouts with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of workouts to return", func() {
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
		Result(ArrayOf(Workout))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a workout")
		Payload(UpdateWorkoutPayload)
		Result(Workout)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a workout")
		Payload(func() {
			Attribute("id", String, "Workout ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
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
