// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen be/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Get    goa.Endpoint
	List   goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		Get:    NewGetEndpoint(s),
		List:   NewListEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Delete: NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Get = m(e.Get)
	e.List = m(e.List)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "user".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateUserPayload)
		return s.Create(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "user".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
		return s.Get(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "user".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "user".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "user".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		return nil, s.Delete(ctx, p)
	}
}
