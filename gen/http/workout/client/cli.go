// Code generated by goa v3.19.1, DO NOT EDIT.
//
// workout HTTP client CLI support package
//
// Command:
// $ goa gen be/design

package client

import (
	workout "be/gen/workout"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the workout create endpoint from
// CLI flags.
func BuildCreatePayload(workoutCreateBody string) (*workout.CreateWorkoutPayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(workoutCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Workout A\",\n      \"trainingPlanId\": \"f47ac10b-58cc-4372-a567-0e02b2c3d479\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.trainingPlanId", body.TrainingPlanID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workout.CreateWorkoutPayload{
		Name:           body.Name,
		TrainingPlanID: body.TrainingPlanID,
	}

	return v, nil
}

// BuildGetPayload builds the payload for the workout get endpoint from CLI
// flags.
func BuildGetPayload(workoutGetID string) (*workout.GetPayload, error) {
	var err error
	var id string
	{
		id = workoutGetID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workout.GetPayload{}
	v.ID = id

	return v, nil
}

// BuildListPayload builds the payload for the workout list endpoint from CLI
// flags.
func BuildListPayload(workoutListLimit string, workoutListOffset string) (*workout.ListPayload, error) {
	var err error
	var limit int
	{
		if workoutListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(workoutListLimit, 10, strconv.IntSize)
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
		if workoutListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(workoutListOffset, 10, strconv.IntSize)
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
	v := &workout.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v, nil
}

// BuildUpdatePayload builds the payload for the workout update endpoint from
// CLI flags.
func BuildUpdatePayload(workoutUpdateBody string, workoutUpdateID string) (*workout.UpdateWorkoutPayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(workoutUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Workout A Updated\",\n      \"trainingPlanId\": \"f47ac10b-58cc-4372-a567-0e02b2c3d479\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.trainingPlanId", body.TrainingPlanID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = workoutUpdateID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workout.UpdateWorkoutPayload{
		Name:           body.Name,
		TrainingPlanID: body.TrainingPlanID,
	}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the workout delete endpoint from
// CLI flags.
func BuildDeletePayload(workoutDeleteID string) (*workout.DeletePayload, error) {
	var err error
	var id string
	{
		id = workoutDeleteID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
	}
	v := &workout.DeletePayload{}
	v.ID = id

	return v, nil
}
