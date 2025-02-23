package workout

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	workoutService "be/gen/workout"

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

func (s *Service) Create(ctx context.Context, payload *workoutService.CreateWorkoutPayload) (*workoutService.Workout, error) {
	w := models.Workout{
		Name:           payload.Name,
		TrainingPlanID: uuid.MustParse(payload.TrainingPlanID),
	}
	saved, err := s.Repository.SaveWorkout(ctx, w)
	if err != nil {
		return nil, err
	}
	return toWorkoutResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *workoutService.GetPayload) (*workoutService.Workout, error) {
	w, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &workoutService.NotFound{Message: "Workout not found"}
		}
		return nil, err
	}
	return toWorkoutResponse(w), nil
}

func (s *Service) List(ctx context.Context, payload *workoutService.ListPayload) ([]*workoutService.Workout, error) {
	workouts, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*workoutService.Workout
	for _, w := range workouts {
		response = append(response, toWorkoutResponse(&w))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *workoutService.UpdateWorkoutPayload) (*workoutService.Workout, error) {
	w, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &workoutService.NotFound{Message: "Workout not found"}
		}
		return nil, err
	}
	w.Name = payload.Name
	w.TrainingPlanID = uuid.MustParse(payload.TrainingPlanID)
	_, err = s.Repository.SaveWorkout(ctx, *w)
	if err != nil {
		return nil, err
	}
	return toWorkoutResponse(w), nil
}

func (s *Service) Delete(ctx context.Context, payload *workoutService.DeletePayload) error {
	w, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &workoutService.NotFound{Message: "Workout not found"}
		}
		return err
	}
	return s.Repository.DeleteWorkout(ctx, w.ID.String())
}

func toWorkoutResponse(w *models.Workout) *workoutService.Workout {
	return &workoutService.Workout{
		ID:             w.ID.String(),
		Name:           w.Name,
		TrainingPlanID: w.TrainingPlanID.String(),
	}
}
