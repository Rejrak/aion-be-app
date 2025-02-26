// Code generated by goa v3.19.1, DO NOT EDIT.
//
// workoutprogress HTTP client CLI support package
//
// Command:
// $ goa gen be/design

package client

import (
	workoutprogress "be/gen/workoutprogress"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the workoutprogress create
// endpoint from CLI flags.
func BuildCreatePayload(workoutprogressCreateBody string) (*workoutprogress.CreateWorkoutProgressPayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(workoutprogressCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date\": \"2025-02-15T10:00:00Z\",\n      \"userId\": \"f47ac10b-58cc-4372-a567-0e02b2c3d479\",\n      \"workoutId\": \"f47ac10b-58cc-4372-a567-0e02b2c3d481\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.workoutId", body.WorkoutID, goa.FormatUUID))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.userId", body.UserID, goa.FormatUUID))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", body.Date, goa.FormatDateTime))
		if err != nil {
			return nil, err
		}
	}
	v := &workoutprogress.CreateWorkoutProgressPayload{
		WorkoutID: body.WorkoutID,
		UserID:    body.UserID,
		Date:      body.Date,
	}

	return v, nil
}

// BuildGetPayload builds the payload for the workoutprogress get endpoint from
// CLI flags.
func BuildGetPayload(workoutprogressGetID string) (*workoutprogress.GetPayload, error) {
	var err error
	var id string
	{
		id = workoutprogressGetID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workoutprogress.GetPayload{}
	v.ID = id

	return v, nil
}

// BuildListPayload builds the payload for the workoutprogress list endpoint
// from CLI flags.
func BuildListPayload(workoutprogressListLimit string, workoutprogressListOffset string) (*workoutprogress.ListPayload, error) {
	var err error
	var limit int
	{
		if workoutprogressListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(workoutprogressListLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var offset int
	{
		if workoutprogressListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(workoutprogressListOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &workoutprogress.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v, nil
}

// BuildUpdatePayload builds the payload for the workoutprogress update
// endpoint from CLI flags.
func BuildUpdatePayload(workoutprogressUpdateBody string, workoutprogressUpdateID string) (*workoutprogress.UpdateWorkoutProgressPayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(workoutprogressUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date\": \"2025-02-15T10:00:00Z\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", body.Date, goa.FormatDateTime))
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = workoutprogressUpdateID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workoutprogress.UpdateWorkoutProgressPayload{
		Date: body.Date,
	}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the workoutprogress delete
// endpoint from CLI flags.
func BuildDeletePayload(workoutprogressDeleteID string) (*workoutprogress.DeletePayload, error) {
	var err error
	var id string
	{
		id = workoutprogressDeleteID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workoutprogress.DeletePayload{}
	v.ID = id

	return v, nil
}
