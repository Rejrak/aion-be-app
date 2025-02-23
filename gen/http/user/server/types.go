// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen be/design

package server

import (
	user "be/gen/user"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Keycloak ID
	KcID *string `form:"kcId,omitempty" json:"kcId,omitempty" xml:"kcId,omitempty"`
	// First name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is admin?
	Admin *bool `form:"admin,omitempty" json:"admin,omitempty" xml:"admin,omitempty"`
}

// UpdateRequestBody is the type of the "user" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Keycloak ID
	KcID *string `form:"kcId,omitempty" json:"kcId,omitempty" xml:"kcId,omitempty"`
	// First name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is admin?
	Admin *bool `form:"admin,omitempty" json:"admin,omitempty" xml:"admin,omitempty"`
}

// CreateResponseBody is the type of the "user" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// Unique ID of the user
	ID string `form:"id" json:"id" xml:"id"`
	// Keycloak ID
	KcID string `form:"kcId" json:"kcId" xml:"kcId"`
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is the user an admin?
	Admin bool `form:"admin" json:"admin" xml:"admin"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Unique ID of the user
	ID string `form:"id" json:"id" xml:"id"`
	// Keycloak ID
	KcID string `form:"kcId" json:"kcId" xml:"kcId"`
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is the user an admin?
	Admin bool `form:"admin" json:"admin" xml:"admin"`
}

// ListResponseBody is the type of the "user" service "list" endpoint HTTP
// response body.
type ListResponseBody []*UserResponse

// UpdateResponseBody is the type of the "user" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// Unique ID of the user
	ID string `form:"id" json:"id" xml:"id"`
	// Keycloak ID
	KcID string `form:"kcId" json:"kcId" xml:"kcId"`
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is the user an admin?
	Admin bool `form:"admin" json:"admin" xml:"admin"`
}

// GetBadRequestResponseBody is the type of the "user" service "get" endpoint
// HTTP response body for the "badRequest" error.
type GetBadRequestResponseBody struct {
	// Nome dell'errore
	Name string `form:"name" json:"name" xml:"name"`
	// ID dell'errore
	ID string `form:"id" json:"id" xml:"id"`
	// Descrizione dettagliata dell'errore
	Message string `form:"message" json:"message" xml:"message"`
	// Indica se l'errore è temporaneo
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Indica se l'errore è dovuto a un timeout
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Indica se l'errore è dovuto a un problema del server
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInternalServerErrorResponseBody is the type of the "user" service "get"
// endpoint HTTP response body for the "internalServerError" error.
type GetInternalServerErrorResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// GetNotFoundResponseBody is the type of the "user" service "get" endpoint
// HTTP response body for the "notFound" error.
type GetNotFoundResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// GetUnauthorizedResponseBody is the type of the "user" service "get" endpoint
// HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateBadRequestResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "badRequest" error.
type UpdateBadRequestResponseBody struct {
	// Nome dell'errore
	Name string `form:"name" json:"name" xml:"name"`
	// ID dell'errore
	ID string `form:"id" json:"id" xml:"id"`
	// Descrizione dettagliata dell'errore
	Message string `form:"message" json:"message" xml:"message"`
	// Indica se l'errore è temporaneo
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Indica se l'errore è dovuto a un timeout
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Indica se l'errore è dovuto a un problema del server
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInternalServerErrorResponseBody is the type of the "user" service
// "update" endpoint HTTP response body for the "internalServerError" error.
type UpdateInternalServerErrorResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateNotFoundResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "notFound" error.
type UpdateNotFoundResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateUnauthorizedResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteBadRequestResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "badRequest" error.
type DeleteBadRequestResponseBody struct {
	// Nome dell'errore
	Name string `form:"name" json:"name" xml:"name"`
	// ID dell'errore
	ID string `form:"id" json:"id" xml:"id"`
	// Descrizione dettagliata dell'errore
	Message string `form:"message" json:"message" xml:"message"`
	// Indica se l'errore è temporaneo
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Indica se l'errore è dovuto a un timeout
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Indica se l'errore è dovuto a un problema del server
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteInternalServerErrorResponseBody is the type of the "user" service
// "delete" endpoint HTTP response body for the "internalServerError" error.
type DeleteInternalServerErrorResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteNotFoundResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "notFound" error.
type DeleteNotFoundResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteUnauthorizedResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "unauthorized" error.
type DeleteUnauthorizedResponseBody struct {
	// Descrizione dell'errore
	Message string `form:"message" json:"message" xml:"message"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// Unique ID of the user
	ID string `form:"id" json:"id" xml:"id"`
	// Keycloak ID
	KcID string `form:"kcId" json:"kcId" xml:"kcId"`
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Nickname
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// Is the user an admin?
	Admin bool `form:"admin" json:"admin" xml:"admin"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "user" service.
func NewCreateResponseBody(res *user.User) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        res.ID,
		KcID:      res.KcID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Nickname:  res.Nickname,
		Admin:     res.Admin,
	}
	{
		var zero bool
		if body.Admin == zero {
			body.Admin = false
		}
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "user" service.
func NewGetResponseBody(res *user.User) *GetResponseBody {
	body := &GetResponseBody{
		ID:        res.ID,
		KcID:      res.KcID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Nickname:  res.Nickname,
		Admin:     res.Admin,
	}
	{
		var zero bool
		if body.Admin == zero {
			body.Admin = false
		}
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "user" service.
func NewListResponseBody(res []*user.User) ListResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUserUserToUserResponse(val)
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "user" service.
func NewUpdateResponseBody(res *user.User) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        res.ID,
		KcID:      res.KcID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Nickname:  res.Nickname,
		Admin:     res.Admin,
	}
	{
		var zero bool
		if body.Admin == zero {
			body.Admin = false
		}
	}
	return body
}

// NewGetBadRequestResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "user" service.
func NewGetBadRequestResponseBody(res *user.BadRequest) *GetBadRequestResponseBody {
	body := &GetBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInternalServerErrorResponseBody builds the HTTP response body from the
// result of the "get" endpoint of the "user" service.
func NewGetInternalServerErrorResponseBody(res *user.InternalServerError) *GetInternalServerErrorResponseBody {
	body := &GetInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "user" service.
func NewGetNotFoundResponseBody(res *user.NotFound) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetUnauthorizedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "user" service.
func NewGetUnauthorizedResponseBody(res *user.Unauthorized) *GetUnauthorizedResponseBody {
	body := &GetUnauthorizedResponseBody{
		Message: res.Message,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "user" service.
func NewUpdateBadRequestResponseBody(res *user.BadRequest) *UpdateBadRequestResponseBody {
	body := &UpdateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "update" endpoint of the "user" service.
func NewUpdateInternalServerErrorResponseBody(res *user.InternalServerError) *UpdateInternalServerErrorResponseBody {
	body := &UpdateInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "user" service.
func NewUpdateNotFoundResponseBody(res *user.NotFound) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Message: res.Message,
	}
	return body
}

// NewUpdateUnauthorizedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "user" service.
func NewUpdateUnauthorizedResponseBody(res *user.Unauthorized) *UpdateUnauthorizedResponseBody {
	body := &UpdateUnauthorizedResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteBadRequestResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "user" service.
func NewDeleteBadRequestResponseBody(res *user.BadRequest) *DeleteBadRequestResponseBody {
	body := &DeleteBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "delete" endpoint of the "user" service.
func NewDeleteInternalServerErrorResponseBody(res *user.InternalServerError) *DeleteInternalServerErrorResponseBody {
	body := &DeleteInternalServerErrorResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "user" service.
func NewDeleteNotFoundResponseBody(res *user.NotFound) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteUnauthorizedResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "user" service.
func NewDeleteUnauthorizedResponseBody(res *user.Unauthorized) *DeleteUnauthorizedResponseBody {
	body := &DeleteUnauthorizedResponseBody{
		Message: res.Message,
	}
	return body
}

// NewCreateUserPayload builds a user service create endpoint payload.
func NewCreateUserPayload(body *CreateRequestBody) *user.CreateUserPayload {
	v := &user.CreateUserPayload{
		KcID:      *body.KcID,
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
		Nickname:  body.Nickname,
	}
	if body.Admin != nil {
		v.Admin = *body.Admin
	}
	if body.Admin == nil {
		v.Admin = false
	}

	return v
}

// NewGetPayload builds a user service get endpoint payload.
func NewGetPayload(id string) *user.GetPayload {
	v := &user.GetPayload{}
	v.ID = id

	return v
}

// NewListPayload builds a user service list endpoint payload.
func NewListPayload(limit int, offset int) *user.ListPayload {
	v := &user.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewUpdatePayload builds a user service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *user.UpdatePayload {
	v := &user.UpdatePayload{
		KcID:      *body.KcID,
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
		Nickname:  body.Nickname,
	}
	if body.Admin != nil {
		v.Admin = *body.Admin
	}
	if body.Admin == nil {
		v.Admin = false
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a user service delete endpoint payload.
func NewDeletePayload(id string) *user.DeletePayload {
	v := &user.DeletePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.KcID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kcId", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstName", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastName", "body"))
	}
	if body.KcID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.kcId", *body.KcID, goa.FormatUUID))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.KcID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kcId", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstName", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastName", "body"))
	}
	if body.KcID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.kcId", *body.KcID, goa.FormatUUID))
	}
	return
}
