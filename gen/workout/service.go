// Code generated by goa v3.19.1, DO NOT EDIT.
//
// workout service
//
// Command:
// $ goa gen be/design

package workout

import (
	"context"
)

// Service for managing workouts
type Service interface {
	// Create a new workout
	Create(context.Context, *CreateWorkoutPayload) (res *Workout, err error)
	// Get a workout by ID
	Get(context.Context, *GetPayload) (res *Workout, err error)
	// List all workouts with pagination
	List(context.Context, *ListPayload) (res []*Workout, err error)
	// Update a workout
	Update(context.Context, *UpdateWorkoutPayload) (res *Workout, err error)
	// Delete a workout
	Delete(context.Context, *DeletePayload) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "aion_service"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "workout"

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

// CreateWorkoutPayload is the payload type of the workout service create
// method.
type CreateWorkoutPayload struct {
	// Name of the workout
	Name string
	// Training Plan ID
	TrainingPlanID string
}

// DeletePayload is the payload type of the workout service delete method.
type DeletePayload struct {
	// Workout ID
	ID string
}

// GetPayload is the payload type of the workout service get method.
type GetPayload struct {
	// Workout ID
	ID string
}

// Errore nel server
type InternalServerError struct {
	// Descrizione dell'errore
	Message string
}

// ListPayload is the payload type of the workout service list method.
type ListPayload struct {
	// Number of workouts to return
	Limit int
	// Offset for pagination
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

// UpdateWorkoutPayload is the payload type of the workout service update
// method.
type UpdateWorkoutPayload struct {
	// Workout ID
	ID string
	// Name of the workout
	Name string
	// Training Plan ID
	TrainingPlanID string
}

// Workout is the result type of the workout service create method.
type Workout struct {
	// Unique ID of the workout
	ID string
	// Name of the workout
	Name string
	// Training Plan ID
	TrainingPlanID string
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
