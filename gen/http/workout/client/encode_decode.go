// Code generated by goa v3.19.1, DO NOT EDIT.
//
// workout HTTP client encoders and decoders
//
// Command:
// $ goa gen be/design

package client

import (
	workout "be/gen/workout"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "workout" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateWorkoutPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("workout", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the workout
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*workout.CreateWorkoutPayload)
		if !ok {
			return goahttp.ErrInvalidType("workout", "create", "*workout.CreateWorkoutPayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("workout", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the workout
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("workout", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("workout", "create", err)
			}
			res := NewCreateWorkoutCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("workout", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "workout" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*workout.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("workout", "get", "*workout.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetWorkoutPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("workout", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the workout
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("workout", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("workout", "get", err)
			}
			res := NewGetWorkoutOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("workout", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "workout" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListWorkoutPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("workout", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the workout list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*workout.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("workout", "list", "*workout.ListPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the workout
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("workout", "list", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateWorkoutResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("workout", "list", err)
			}
			res := NewListWorkoutOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("workout", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "workout" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*workout.UpdateWorkoutPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("workout", "update", "*workout.UpdateWorkoutPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateWorkoutPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("workout", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the workout
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*workout.UpdateWorkoutPayload)
		if !ok {
			return goahttp.ErrInvalidType("workout", "update", "*workout.UpdateWorkoutPayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("workout", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the workout
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("workout", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("workout", "update", err)
			}
			res := NewUpdateWorkoutOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("workout", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "workout" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*workout.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("workout", "delete", "*workout.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteWorkoutPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("workout", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteResponse returns a decoder for responses returned by the workout
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("workout", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalWorkoutResponseToWorkoutWorkout builds a value of type
// *workout.Workout from a value of type *WorkoutResponse.
func unmarshalWorkoutResponseToWorkoutWorkout(v *WorkoutResponse) *workout.Workout {
	res := &workout.Workout{
		ID:             *v.ID,
		Name:           *v.Name,
		TrainingPlanID: *v.TrainingPlanID,
	}

	return res
}
