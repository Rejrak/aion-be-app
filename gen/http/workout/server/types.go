// Code generated by goa v3.19.1, DO NOT EDIT.
//
// workout HTTP server types
//
// Command:
// $ goa gen be/design

package server

import (
	workout "be/gen/workout"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "workout" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name of the workout
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Training Plan ID
	TrainingPlanID *string `form:"trainingPlanId,omitempty" json:"trainingPlanId,omitempty" xml:"trainingPlanId,omitempty"`
}

// UpdateRequestBody is the type of the "workout" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name of the workout
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Training Plan ID
	TrainingPlanID *string `form:"trainingPlanId,omitempty" json:"trainingPlanId,omitempty" xml:"trainingPlanId,omitempty"`
}

// CreateResponseBody is the type of the "workout" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// Unique ID of the workout
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the workout
	Name string `form:"name" json:"name" xml:"name"`
	// Training Plan ID
	TrainingPlanID string `form:"trainingPlanId" json:"trainingPlanId" xml:"trainingPlanId"`
}

// GetResponseBody is the type of the "workout" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Unique ID of the workout
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the workout
	Name string `form:"name" json:"name" xml:"name"`
	// Training Plan ID
	TrainingPlanID string `form:"trainingPlanId" json:"trainingPlanId" xml:"trainingPlanId"`
}

// ListResponseBody is the type of the "workout" service "list" endpoint HTTP
// response body.
type ListResponseBody []*WorkoutResponse

// UpdateResponseBody is the type of the "workout" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// Unique ID of the workout
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the workout
	Name string `form:"name" json:"name" xml:"name"`
	// Training Plan ID
	TrainingPlanID string `form:"trainingPlanId" json:"trainingPlanId" xml:"trainingPlanId"`
}

// WorkoutResponse is used to define fields on response body types.
type WorkoutResponse struct {
	// Unique ID of the workout
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the workout
	Name string `form:"name" json:"name" xml:"name"`
	// Training Plan ID
	TrainingPlanID string `form:"trainingPlanId" json:"trainingPlanId" xml:"trainingPlanId"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "workout" service.
func NewCreateResponseBody(res *workout.Workout) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:             res.ID,
		Name:           res.Name,
		TrainingPlanID: res.TrainingPlanID,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "workout" service.
func NewGetResponseBody(res *workout.Workout) *GetResponseBody {
	body := &GetResponseBody{
		ID:             res.ID,
		Name:           res.Name,
		TrainingPlanID: res.TrainingPlanID,
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "workout" service.
func NewListResponseBody(res []*workout.Workout) ListResponseBody {
	body := make([]*WorkoutResponse, len(res))
	for i, val := range res {
		body[i] = marshalWorkoutWorkoutToWorkoutResponse(val)
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "workout" service.
func NewUpdateResponseBody(res *workout.Workout) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:             res.ID,
		Name:           res.Name,
		TrainingPlanID: res.TrainingPlanID,
	}
	return body
}

// NewCreateWorkoutPayload builds a workout service create endpoint payload.
func NewCreateWorkoutPayload(body *CreateRequestBody) *workout.CreateWorkoutPayload {
	v := &workout.CreateWorkoutPayload{
		Name:           *body.Name,
		TrainingPlanID: *body.TrainingPlanID,
	}

	return v
}

// NewGetPayload builds a workout service get endpoint payload.
func NewGetPayload(id string) *workout.GetPayload {
	v := &workout.GetPayload{}
	v.ID = id

	return v
}

// NewListPayload builds a workout service list endpoint payload.
func NewListPayload(limit int, offset int) *workout.ListPayload {
	v := &workout.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewUpdateWorkoutPayload builds a workout service update endpoint payload.
func NewUpdateWorkoutPayload(body *UpdateRequestBody, id string) *workout.UpdateWorkoutPayload {
	v := &workout.UpdateWorkoutPayload{
		Name:           *body.Name,
		TrainingPlanID: *body.TrainingPlanID,
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a workout service delete endpoint payload.
func NewDeletePayload(id string) *workout.DeletePayload {
	v := &workout.DeletePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.TrainingPlanID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("trainingPlanId", "body"))
	}
	if body.TrainingPlanID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.trainingPlanId", *body.TrainingPlanID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.TrainingPlanID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("trainingPlanId", "body"))
	}
	if body.TrainingPlanID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.trainingPlanId", *body.TrainingPlanID, goa.FormatUUID))
	}
	return
}
