package services

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var TrainingPlan = Type("TrainingPlan", func() {
	Attribute("id", String, "Unique ID of the training plan", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the training plan", func() {
		Example("Piano di Allenamento")
	})
	Attribute("description", String, "Description", func() {
		Example("Descrizione del piano di allenamento")
	})
	Attribute("startDate", String, "Start date", func() {
		Example("2025-02-01T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("endDate", String, "End date", func() {
		Example("2025-02-28T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("userId", String, "User ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("workoutTypeId", String, "Workout Type ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d480")
		Format(FormatUUID)
	})
	Required("id", "name", "startDate", "endDate", "userId", "workoutTypeId")
})

var CreateTrainingPlanPayload = Type("CreateTrainingPlanPayload", func() {
	Attribute("name", String, "Name of the training plan", func() {
		Example("Piano di Allenamento")
	})
	Attribute("description", String, "Description", func() {
		Example("Descrizione del piano di allenamento")
	})
	Attribute("startDate", String, "Start date", func() {
		Example("2025-02-01T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("endDate", String, "End date", func() {
		Example("2025-02-28T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("userId", String, "User ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("workoutTypeId", String, "Workout Type ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d480")
		Format(FormatUUID)
	})
	Required("name", "startDate", "endDate", "userId", "workoutTypeId")
})

var UpdateTrainingPlanPayload = Type("UpdateTrainingPlanPayload", func() {
	Attribute("id", String, "Unique ID of the training plan", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("name", String, "Name of the training plan", func() {
		Example("Piano di Allenamento Updated")
	})
	Attribute("description", String, "Description", func() {
		Example("Descrizione aggiornata del piano")
	})
	Attribute("startDate", String, "Start date", func() {
		Example("2025-02-01T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("endDate", String, "End date", func() {
		Example("2025-02-28T00:00:00Z")
		Format(FormatDateTime)
	})
	Attribute("userId", String, "User ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		Format(FormatUUID)
	})
	Attribute("workoutTypeId", String, "Workout Type ID", func() {
		Example("f47ac10b-58cc-4372-a567-0e02b2c3d480")
		Format(FormatUUID)
	})
	Required("id", "name", "startDate", "endDate", "userId", "workoutTypeId")
})

var TrainingPlanService = Service("trainingplan", func() {
	Description("Service for managing training plans")

	HTTP(func() {
		Path("/trainingplans")
	})
	Error("unauthorized", errors.Unauthorized, "Unauthorized")
	Error("internalServerError", errors.InternalServerError, "Internal server error")
	Error("notFound", errors.NotFound, "Training plan not found")
	Error("badRequest", errors.BadRequest, "Invalid request")

	Method("create", func() {
		Description("Create a new training plan")
		Payload(CreateTrainingPlanPayload)
		Result(TrainingPlan)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Get a training plan by ID")
		Payload(func() {
			Attribute("id", String, "Training plan ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
				Format(FormatUUID)
			})
			Required("id")
		})
		Result(TrainingPlan)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("List all training plans with pagination")
		Payload(func() {
			Attribute("limit", Int, "Number of training plans to return", func() {
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
		Result(ArrayOf(TrainingPlan))
		HTTP(func() {
			GET("")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a training plan")
		Payload(UpdateTrainingPlanPayload)
		Result(TrainingPlan)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a training plan")
		Payload(func() {
			Attribute("id", String, "Training plan ID", func() {
				Example("f47ac10b-58cc-4372-a567-0e02b2c3d479")
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
