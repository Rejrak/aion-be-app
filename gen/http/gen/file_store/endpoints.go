// Code generated by goa v3.19.1, DO NOT EDIT.
//
// FileStore endpoints
//
// Command:
// $ goa gen aion/design

package filestore

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "FileStore" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "FileStore" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "FileStore" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}
