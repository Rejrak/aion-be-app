package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var Exercise = Type("Exercise", func() {
	Attribute("id", String, "Unique ID of the exercise", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the exercise", func() {
		Example("Push Up")
	})
	Attribute("muscleGroupId", String, "Muscle Group ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
		Format(FormatUUID)
	})
	Required("id", "name", "muscleGroupId")
})

var CreateExercisePayload = Type("CreateExercisePayload", func() {
	Attribute("name", String, "Name of the exercise", func() {
		Example("Push Up")
	})
	Attribute("muscleGroupId", String, "Muscle Group ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
		Format(FormatUUID)
	})
	Required("name", "muscleGroupId")
})

var UpdateExercisePayload = Type("UpdateExercisePayload", func() {
	Attribute("id", String, "Exercise ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the exercise", func() {
		Example("Push Up Updated")
	})
	Attribute("muscleGroupId", String, "Muscle Group ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
		Format(FormatUUID)
	})
	Required("id", "name", "muscleGroupId")
})

var ExerciseService = Service("exercise", func() {
	Description("Service for managing exercises")

	HTTP(func() {
		Path("/exercises")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Exercise not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new exercise")
		Payload(CreateExercisePayload)
		Result(Exercise)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get an exercise by ID")
		Payload(func() {
			Attribute("id", String, "Exercise ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(Exercise)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all exercises with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of exercises to return", func() {
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
		Result(ArrayOf(Exercise))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update an exercise")
		Payload(UpdateExercisePayload)
		Result(Exercise)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete an exercise")
		Payload(func() {
			Attribute("id", String, "Exercise ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d484")
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
