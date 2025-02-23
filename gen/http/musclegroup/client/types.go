// Code generated by goa v3.19.1, DO NOT EDIT.
//
// musclegroup HTTP client types
//
// Command:
// $ goa gen be/design

package client

import (
	musclegroup "be/gen/musclegroup"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "musclegroup" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// UpdateRequestBody is the type of the "musclegroup" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name of the muscle group
	Name string `form:"name" json:"name" xml:"name"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// CreateResponseBody is the type of the "musclegroup" service "create"
// endpoint HTTP response body.
type CreateResponseBody struct {
	// Unique ID of the muscle group
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// GetResponseBody is the type of the "musclegroup" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Unique ID of the muscle group
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
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
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// MuscleGroupResponse is used to define fields on response body types.
type MuscleGroupResponse struct {
	// Unique ID of the muscle group
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the muscle group
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "musclegroup" service.
func NewCreateRequestBody(p *musclegroup.CreateMuscleGroupPayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:        p.Name,
		Description: p.Description,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "musclegroup" service.
func NewUpdateRequestBody(p *musclegroup.UpdateMuscleGroupPayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name:        p.Name,
		Description: p.Description,
	}
	return body
}

// NewCreateMuscleGroupCreated builds a "musclegroup" service "create" endpoint
// result from a HTTP "Created" response.
func NewCreateMuscleGroupCreated(body *CreateResponseBody) *musclegroup.MuscleGroup {
	v := &musclegroup.MuscleGroup{
		ID:          *body.ID,
		Name:        *body.Name,
		Description: body.Description,
	}

	return v
}

// NewGetMuscleGroupOK builds a "musclegroup" service "get" endpoint result
// from a HTTP "OK" response.
func NewGetMuscleGroupOK(body *GetResponseBody) *musclegroup.MuscleGroup {
	v := &musclegroup.MuscleGroup{
		ID:          *body.ID,
		Name:        *body.Name,
		Description: body.Description,
	}

	return v
}

// NewListMuscleGroupOK builds a "musclegroup" service "list" endpoint result
// from a HTTP "OK" response.
func NewListMuscleGroupOK(body []*MuscleGroupResponse) []*musclegroup.MuscleGroup {
	v := make([]*musclegroup.MuscleGroup, len(body))
	for i, val := range body {
		v[i] = unmarshalMuscleGroupResponseToMusclegroupMuscleGroup(val)
	}

	return v
}

// NewUpdateMuscleGroupOK builds a "musclegroup" service "update" endpoint
// result from a HTTP "OK" response.
func NewUpdateMuscleGroupOK(body *UpdateResponseBody) *musclegroup.MuscleGroup {
	v := &musclegroup.MuscleGroup{
		ID:          *body.ID,
		Name:        *body.Name,
		Description: body.Description,
	}

	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}

// ValidateMuscleGroupResponse runs the validations defined on
// MuscleGroupResponse
func ValidateMuscleGroupResponse(body *MuscleGroupResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	return
}
