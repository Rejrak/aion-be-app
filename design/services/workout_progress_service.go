package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var WorkoutProgress = Type("WorkoutProgress", func() {
	Attribute("id", String, "Unique ID of the workout progress", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
		Format(FormatUUID)
	})
	Attribute("workoutId", String, "Workout ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("userId", String, "User ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("date", String, "Date of the workout progress", func() {
		Example("2025-02-15T10:00:00Z")
		Format(FormatDateTime)
	})
	Required("id", "workoutId", "userId", "date")
})

var CreateWorkoutProgressPayload = Type("CreateWorkoutProgressPayload", func() {
	Attribute("workoutId", String, "Workout ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("userId", String, "User ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("date", String, "Date of the workout progress", func() {
		Example("2025-02-15T10:00:00Z")
		Format(FormatDateTime)
	})
	Required("workoutId", "userId", "date")
})

var UpdateWorkoutProgressPayload = Type("UpdateWorkoutProgressPayload", func() {
	Attribute("id", String, "Workout Progress ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
		Format(FormatUUID)
	})
	Attribute("date", String, "Date of the workout progress", func() {
		Example("2025-02-15T10:00:00Z")
		Format(FormatDateTime)
	})
	Required("id", "date")
})

var WorkoutProgressService = Service("workoutprogress", func() {
	Description("Service for managing workout progress")

	HTTP(func() {
		Path("/workoutprogress")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Workout progress not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new workout progress record")
		Payload(CreateWorkoutProgressPayload)
		Result(WorkoutProgress)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a workout progress record by ID")
		Payload(func() {
			Attribute("id", String, "Workout Progress ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(WorkoutProgress)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all workout progress records with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of records to return", func() {
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
		Result(ArrayOf(WorkoutProgress))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a workout progress record")
		Payload(UpdateWorkoutProgressPayload)
		Result(WorkoutProgress)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a workout progress record")
		Payload(func() {
			Attribute("id", String, "Workout Progress ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
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
