package workouttype

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	workoutTypeService "be/gen/workouttype"

	"gorm.io/gorm"
)

type Service struct {
	Repository *Repository
}

func NewService() *Service {
	return &Service{
		Repository: NewRepository(db.DB.Database),
	}
}

func (s *Service) Create(ctx context.Context, payload *workoutTypeService.CreateWorkoutTypePayload) (*workoutTypeService.WorkoutType, error) {
	wt := models.WorkoutType{
		Name:        payload.Name,
		Description: *payload.Description,
	}
	saved, err := s.Repository.SaveWorkoutType(ctx, wt)
	if err != nil {
		return nil, err
	}
	return toWorkoutTypeResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *workoutTypeService.GetPayload) (*workoutTypeService.WorkoutType, error) {
	wt, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &workoutTypeService.NotFound{Message: "Workout type not found"}
		}
		return nil, err
	}
	return toWorkoutTypeResponse(wt), nil
}

func (s *Service) List(ctx context.Context, payload *workoutTypeService.ListPayload) ([]*workoutTypeService.WorkoutType, error) {
	wts, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*workoutTypeService.WorkoutType
	for _, wt := range wts {
		response = append(response, toWorkoutTypeResponse(&wt))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *workoutTypeService.UpdateWorkoutTypePayload) (*workoutTypeService.WorkoutType, error) {
	wt, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &workoutTypeService.NotFound{Message: "Workout type not found"}
		}
		return nil, err
	}
	wt.Name = payload.Name
	wt.Description = *payload.Description
	_, err = s.Repository.SaveWorkoutType(ctx, *wt)
	if err != nil {
		return nil, err
	}
	return toWorkoutTypeResponse(wt), nil
}

func (s *Service) Delete(ctx context.Context, payload *workoutTypeService.DeletePayload) error {
	wt, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &workoutTypeService.NotFound{Message: "Workout type not found"}
		}
		return err
	}
	return s.Repository.DeleteWorkoutType(ctx, wt.ID.String())
}

func toWorkoutTypeResponse(wt *models.WorkoutType) *workoutTypeService.WorkoutType {
	return &workoutTypeService.WorkoutType{
		ID:          wt.ID.String(),
		Name:        wt.Name,
		Description: &wt.Description,
	}
}
