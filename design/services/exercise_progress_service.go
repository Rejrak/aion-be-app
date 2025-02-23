package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var ExerciseProgress = Type("ExerciseProgress", func() {
	Attribute("id", String, "Unique ID of the exercise progress", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d487")
		Format(FormatUUID)
	})
	Attribute("workoutProgressId", String, "Workout Progress ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
		Format(FormatUUID)
	})
	Attribute("workoutExerciseId", String, "Workout Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
		Format(FormatUUID)
	})
	Attribute("actualRepetitions", Int, "Actual repetitions performed", func() {
		Example(10)
	})
	Attribute("actualWeight", Float32, "Actual weight used", func() {
		Example(50.0)
	})
	Attribute("actualDuration", Int, "Actual duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Felt strong")
	})
	Required("id", "workoutProgressId", "workoutExerciseId", "actualRepetitions")
})

var CreateExerciseProgressPayload = Type("CreateExerciseProgressPayload", func() {
	Attribute("workoutProgressId", String, "Workout Progress ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d486")
		Format(FormatUUID)
	})
	Attribute("workoutExerciseId", String, "Workout Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
		Format(FormatUUID)
	})
	Attribute("actualRepetitions", Int, "Actual repetitions performed", func() {
		Example(10)
	})
	Attribute("actualWeight", Float32, "Actual weight used", func() {
		Example(50.0)
	})
	Attribute("actualDuration", Int, "Actual duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Felt strong")
	})
	Required("workoutProgressId", "workoutExerciseId", "actualRepetitions")
})

var UpdateExerciseProgressPayload = Type("UpdateExerciseProgressPayload", func() {
	Attribute("id", String, "Exercise Progress ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d487")
		Format(FormatUUID)
	})
	Attribute("actualRepetitions", Int, "Actual repetitions performed", func() {
		Example(10)
	})
	Attribute("actualWeight", Float32, "Actual weight used", func() {
		Example(50.0)
	})
	Attribute("actualDuration", Int, "Actual duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Felt strong")
	})
	Required("id", "actualRepetitions")
})

var ExerciseProgressService = Service("exerciseprogress", func() {
	Description("Service for managing exercise progress")

	HTTP(func() {
		Path("/exerciseprogress")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Exercise progress not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new exercise progress record")
		Payload(CreateExerciseProgressPayload)
		Result(ExerciseProgress)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get an exercise progress record by ID")
		Payload(func() {
			Attribute("id", String, "Exercise Progress ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d487")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(ExerciseProgress)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all exercise progress records with pagination")
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
		Result(ArrayOf(ExerciseProgress))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update an exercise progress record")
		Payload(UpdateExerciseProgressPayload)
		Result(ExerciseProgress)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete an exercise progress record")
		Payload(func() {
			Attribute("id", String, "Exercise Progress ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d487")
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
