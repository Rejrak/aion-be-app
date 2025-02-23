package exercise

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	exerciseService "be/gen/exercise"

	"github.com/google/uuid"
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

func (s *Service) Create(ctx context.Context, payload *exerciseService.CreateExercisePayload) (*exerciseService.Exercise, error) {
	ex := models.Exercise{
		Name:          payload.Name,
		MuscleGroupID: uuid.MustParse(payload.MuscleGroupID),
	}
	saved, err := s.Repository.SaveExercise(ctx, ex)
	if err != nil {
		return nil, err
	}
	return toExerciseResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *exerciseService.GetPayload) (*exerciseService.Exercise, error) {
	ex, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &exerciseService.NotFound{Message: "Exercise not found"}
		}
		return nil, err
	}
	return toExerciseResponse(ex), nil
}

func (s *Service) List(ctx context.Context, payload *exerciseService.ListPayload) ([]*exerciseService.Exercise, error) {
	exs, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*exerciseService.Exercise
	for _, ex := range exs {
		response = append(response, toExerciseResponse(&ex))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *exerciseService.UpdateExercisePayload) (*exerciseService.Exercise, error) {
	ex, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &exerciseService.NotFound{Message: "Exercise not found"}
		}
		return nil, err
	}
	ex.Name = payload.Name
	ex.MuscleGroupID = uuid.MustParse(payload.MuscleGroupID)
	_, err = s.Repository.SaveExercise(ctx, *ex)
	if err != nil {
		return nil, err
	}
	return toExerciseResponse(ex), nil
}

func (s *Service) Delete(ctx context.Context, payload *exerciseService.DeletePayload) error {
	ex, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &exerciseService.NotFound{Message: "Exercise not found"}
		}
		return err
	}
	return s.Repository.DeleteExercise(ctx, ex.ID.String())
}

func toExerciseResponse(ex *models.Exercise) *exerciseService.Exercise {
	return &exerciseService.Exercise{
		ID:            ex.ID.String(),
		Name:          ex.Name,
		MuscleGroupID: ex.MuscleGroupID.String(),
	}
}
