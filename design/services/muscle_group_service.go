package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var MuscleGroup = Type("MuscleGroup", func() {
	Attribute("id", String, "Unique ID of the muscle group", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the muscle group", func() {
		Example("Chest")
	})
	Attribute("description", String, "Description", func() {
		Example("Group for chest exercises")
	})
	Required("id", "name")
})

var CreateMuscleGroupPayload = Type("CreateMuscleGroupPayload", func() {
	Attribute("name", String, "Name of the muscle group", func() {
		Example("Chest")
	})
	Attribute("description", String, "Description", func() {
		Example("Group for chest exercises")
	})
	Required("name")
})

var UpdateMuscleGroupPayload = Type("UpdateMuscleGroupPayload", func() {
	Attribute("id", String, "Muscle Group ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the muscle group", func() {
		Example("Chest Updated")
	})
	Attribute("description", String, "Description", func() {
		Example("Updated description")
	})
	Required("id", "name")
})

var MuscleGroupService = Service("musclegroup", func() {
	Description("Service for managing muscle groups")

	HTTP(func() {
		Path("/musclegroups")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Muscle group not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new muscle group")
		Payload(CreateMuscleGroupPayload)
		Result(MuscleGroup)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a muscle group by ID")
		Payload(func() {
			Attribute("id", String, "Muscle Group ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(MuscleGroup)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all muscle groups with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of muscle groups to return", func() {
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
		Result(ArrayOf(MuscleGroup))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a muscle group")
		Payload(UpdateMuscleGroupPayload)
		Result(MuscleGroup)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a muscle group")
		Payload(func() {
			Attribute("id", String, "Muscle Group ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d483")
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
