// Code generated by goa v3.19.1, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen aion/design

package server

import (
	"fmt"
)

// CreateUserPath returns the URL path to the user service create HTTP endpoint.
func CreateUserPath() string {
	return "/api/v1/user"
}

// GetUserPath returns the URL path to the user service get HTTP endpoint.
func GetUserPath(id string) string {
	return fmt.Sprintf("/api/v1/user/%v", id)
}

// ListUserPath returns the URL path to the user service list HTTP endpoint.
func ListUserPath() string {
	return "/api/v1/user"
}

// UpdateUserPath returns the URL path to the user service update HTTP endpoint.
func UpdateUserPath(id string) string {
	return fmt.Sprintf("/api/v1/user/%v", id)
}

// DeleteUserPath returns the URL path to the user service delete HTTP endpoint.
func DeleteUserPath(id string) string {
	return fmt.Sprintf("/api/v1/user/%v", id)
}
