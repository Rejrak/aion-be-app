// Code generated by goa v3.19.1, DO NOT EDIT.
//
// musclegroup HTTP server types
//
// Command:
// $ goa gen be/design

package server

import (
	musclegroup "be/gen/musclegroup"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "musclegroup" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// UpdateRequestBody is the type of the "musclegroup" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// CreateResponseBody is the type of the "musclegroup" service "create"
// endpoint HTTP response body.
type CreateResponseBody struct {
	// Unique ID of the muscle group
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// GetResponseBody is the type of the "musclegroup" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Unique ID of the muscle group
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// ListResponseBody is the type of the "musclegroup" service "list" endpoint
// HTTP response body.
type ListResponseBody []*MuscleGroupResponse

// UpdateResponseBody is the type of the "musclegroup" service "update"
// endpoint HTTP response body.
type UpdateResponseBody struct {
	// Unique ID of the muscle group
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// MuscleGroupResponse is used to define fields on response body types.
type MuscleGroupResponse struct {
	// Unique ID of the muscle group
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "musclegroup" service.
func NewCreateResponseBody(res *musclegroup.MuscleGroup) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "musclegroup" service.
func NewGetResponseBody(res *musclegroup.MuscleGroup) *GetResponseBody {
	body := &GetResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "musclegroup" service.
func NewListResponseBody(res []*musclegroup.MuscleGroup) ListResponseBody {
	body := make([]*MuscleGroupResponse, len(res))
	for i, val := range res {
		body[i] = marshalMusclegroupMuscleGroupToMuscleGroupResponse(val)
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "musclegroup" service.
func NewUpdateResponseBody(res *musclegroup.MuscleGroup) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
	}
	return body
}

// NewCreateMuscleGroupPayload builds a musclegroup service create endpoint
// payload.
func NewCreateMuscleGroupPayload(body *CreateRequestBody) *musclegroup.CreateMuscleGroupPayload {
	v := &musclegroup.CreateMuscleGroupPayload{
		Name:        *body.Name,
		Description: body.Description,
	}

	return v
}

// NewGetPayload builds a musclegroup service get endpoint payload.
func NewGetPayload(id string) *musclegroup.GetPayload {
	v := &musclegroup.GetPayload{}
	v.ID = id

	return v
}

// NewListPayload builds a musclegroup service list endpoint payload.
func NewListPayload(limit int, offset int) *musclegroup.ListPayload {
	v := &musclegroup.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewUpdateMuscleGroupPayload builds a musclegroup service update endpoint
// payload.
func NewUpdateMuscleGroupPayload(body *UpdateRequestBody, id string) *musclegroup.UpdateMuscleGroupPayload {
	v := &musclegroup.UpdateMuscleGroupPayload{
		Name:        *body.Name,
		Description: body.Description,
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a musclegroup service delete endpoint payload.
func NewDeletePayload(id string) *musclegroup.DeletePayload {
	v := &musclegroup.DeletePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
