package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var WorkoutExercise = Type("WorkoutExercise", func() {
	Attribute("id", String, "Unique ID of the workout exercise", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
		Format(FormatUUID)
	})
	Attribute("workoutId", String, "Workout ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("exerciseId", String, "Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
		Format(FormatUUID)
	})
	Attribute("sets", Int, "Number of sets", func() {
		Example(3)
	})
	Attribute("repetitions", Int, "Number of repetitions", func() {
		Example(12)
	})
	Attribute("duration", Int, "Duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Focus on form")
	})
	Required("id", "workoutId", "exerciseId", "sets", "repetitions")
})

var CreateWorkoutExercisePayload = Type("CreateWorkoutExercisePayload", func() {
	Attribute("workoutId", String, "Workout ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d481")
		Format(FormatUUID)
	})
	Attribute("exerciseId", String, "Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
		Format(FormatUUID)
	})
	Attribute("sets", Int, "Number of sets", func() {
		Example(3)
	})
	Attribute("repetitions", Int, "Number of repetitions", func() {
		Example(12)
	})
	Attribute("duration", Int, "Duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Focus on form")
	})
	Required("workoutId", "exerciseId", "sets", "repetitions")
})

var UpdateWorkoutExercisePayload = Type("UpdateWorkoutExercisePayload", func() {
	Attribute("id", String, "Workout Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
		Format(FormatUUID)
	})
	Attribute("sets", Int, "Number of sets", func() {
		Example(3)
	})
	Attribute("repetitions", Int, "Number of repetitions", func() {
		Example(12)
	})
	Attribute("duration", Int, "Duration in seconds", func() {
		Example(60)
	})
	Attribute("notes", String, "Notes", func() {
		Example("Focus on form")
	})
	Required("id", "sets", "repetitions")
})

var WorkoutExerciseService = Service("workoutexercise", func() {
	Description("Service for managing workout exercise configuration")

	HTTP(func() {
		Path("/workoutexercises")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Workout exercise not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new workout exercise configuration")
		Payload(CreateWorkoutExercisePayload)
		Result(WorkoutExercise)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a workout exercise configuration by ID")
		Payload(func() {
			Attribute("id", String, "Workout Exercise ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(WorkoutExercise)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all workout exercise configurations with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of items to return", func() {
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
		Result(ArrayOf(WorkoutExercise))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a workout exercise configuration")
		Payload(UpdateWorkoutExercisePayload)
		Result(WorkoutExercise)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a workout exercise configuration")
		Payload(func() {
			Attribute("id", String, "Workout Exercise ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d485")
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
