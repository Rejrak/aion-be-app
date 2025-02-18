// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen aion/design

package user

import (
	"context"
)

// User service for managing users
type Service interface {
	// Create a new user
	Create(context.Context, *CreateUserPayload) (res *User, err error)
	// Get a user by ID
	Get(context.Context, *GetPayload) (res *User, err error)
	// List all users with pagination
	List(context.Context, *ListPayload) (res []*User, err error)
	// Update a user
	Update(context.Context, *UpdatePayload) (res *User, err error)
	// Delete a user
	Delete(context.Context, *DeletePayload) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "aion_service"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"create", "get", "list", "update", "delete"}

// Body di risposta per la richiesta non valida (400)
type BadRequest struct {
	// Nome dell'errore
	Name string
	// ID dell'errore
	ID string
	// Descrizione dettagliata dell'errore
	Message string
	// Indica se l'errore è temporaneo
	Temporary bool
	// Indica se l'errore è dovuto a un timeout
	Timeout bool
	// Indica se l'errore è dovuto a un problema del server
	Fault bool
}

// CreateUserPayload is the payload type of the user service create method.
type CreateUserPayload struct {
	// Keycloak ID
	KcID string
	// First name
	FirstName string
	// Last name
	LastName string
	// Nickname
	Nickname *string
	// Is admin?
	Admin bool
}

// DeletePayload is the payload type of the user service delete method.
type DeletePayload struct {
	// User ID
	ID string
}

// GetPayload is the payload type of the user service get method.
type GetPayload struct {
	// User ID
	ID string
}

// Errore nel server
type InternalServerError struct {
	// Descrizione dell'errore
	Message string
}

// ListPayload is the payload type of the user service list method.
type ListPayload struct {
	// Number of users to return per page
	Limit int
	// Number of users to skip
	Offset int
}

// Dato non trovato all'interno del sistema
type NotFound struct {
	// Descrizione dell'errore
	Message string
}

// Errore di autenticazione non riuscita
type Unauthorized struct {
	// Descrizione dell'errore
	Message string
}

// UpdatePayload is the payload type of the user service update method.
type UpdatePayload struct {
	// User ID
	ID string
	// Keycloak ID
	KcID string
	// First name
	FirstName string
	// Last name
	LastName string
	// Nickname
	Nickname *string
	// Is admin?
	Admin bool
}

// User is the result type of the user service create method.
type User struct {
	// Unique ID of the user
	ID string
	// Keycloak ID
	KcID string
	// First name of the user
	FirstName string
	// Last name of the user
	LastName string
	// Nickname
	Nickname *string
	// Is the user an admin?
	Admin bool
}

// Error returns an error description.
func (e *BadRequest) Error() string {
	return "Body di risposta per la richiesta non valida (400)"
}

// ErrorName returns "BadRequest".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *BadRequest) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "BadRequest".
func (e *BadRequest) GoaErrorName() string {
	return "badRequest"
}

// Error returns an error description.
func (e *InternalServerError) Error() string {
	return "Errore nel server"
}

// ErrorName returns "InternalServerError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InternalServerError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InternalServerError".
func (e *InternalServerError) GoaErrorName() string {
	return "internalServerError"
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "Dato non trovato all'interno del sistema "
}

// ErrorName returns "NotFound".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFound".
func (e *NotFound) GoaErrorName() string {
	return "notFound"
}

// Error returns an error description.
func (e *Unauthorized) Error() string {
	return "Errore di autenticazione non riuscita"
}

// ErrorName returns "Unauthorized".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *Unauthorized) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "Unauthorized".
func (e *Unauthorized) GoaErrorName() string {
	return "unauthorized"
}
