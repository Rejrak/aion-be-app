package exerciseprogress

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	epService "be/gen/exerciseprogress"

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

func (s *Service) Create(ctx context.Context, payload *epService.CreateExerciseProgressPayload) (*epService.ExerciseProgress, error) {
	ep := models.ExerciseProgress{
		WorkoutProgressID: uuid.MustParse(payload.WorkoutProgressID),
		WorkoutExerciseID: uuid.MustParse(payload.WorkoutExerciseID),
		ActualRepetitions: payload.ActualRepetitions,
		ActualWeight:      float64(*payload.ActualWeight),
		ActualDuration:    *payload.ActualDuration,
		Notes:             *payload.Notes,
	}
	saved, err := s.Repository.SaveExerciseProgress(ctx, ep)
	if err != nil {
		return nil, err
	}
	return toExerciseProgressResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *epService.GetPayload) (*epService.ExerciseProgress, error) {
	ep, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &epService.NotFound{Message: "Exercise progress not found"}
		}
		return nil, err
	}
	return toExerciseProgressResponse(ep), nil
}

func (s *Service) List(ctx context.Context, payload *epService.ListPayload) ([]*epService.ExerciseProgress, error) {
	eps, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*epService.ExerciseProgress
	for _, ep := range eps {
		response = append(response, toExerciseProgressResponse(&ep))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *epService.UpdateExerciseProgressPayload) (*epService.ExerciseProgress, error) {
	ep, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &epService.NotFound{Message: "Exercise progress not found"}
		}
		return nil, err
	}
	ep.ActualRepetitions = payload.ActualRepetitions
	ep.ActualWeight = float64(*payload.ActualWeight)
	ep.ActualDuration = *payload.ActualDuration
	ep.Notes = *payload.Notes
	_, err = s.Repository.SaveExerciseProgress(ctx, *ep)
	if err != nil {
		return nil, err
	}
	return toExerciseProgressResponse(ep), nil
}

func (s *Service) Delete(ctx context.Context, payload *epService.DeletePayload) error {
	ep, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &epService.NotFound{Message: "Exercise progress not found"}
		}
		return err
	}
	return s.Repository.DeleteExerciseProgress(ctx, ep.ID.String())
}

func toExerciseProgressResponse(ep *models.ExerciseProgress) *epService.ExerciseProgress {
	var actualWeight = float32(ep.ActualWeight)
	var actualDuration = int(ep.ActualDuration)
	return &epService.ExerciseProgress{
		ID:                ep.ID.String(),
		WorkoutProgressID: ep.WorkoutProgressID.String(),
		WorkoutExerciseID: ep.WorkoutExerciseID.String(),
		ActualRepetitions: ep.ActualRepetitions,
		ActualWeight:      &actualWeight,
		ActualDuration:    &actualDuration,
		Notes:             &ep.Notes,
	}
}
