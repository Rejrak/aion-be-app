package workoutexercise

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	weService "be/gen/workoutexercise"

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

func (s *Service) Create(ctx context.Context, payload *weService.CreateWorkoutExercisePayload) (*weService.WorkoutExercise, error) {
	we := models.WorkoutExercise{
		WorkoutID:   uuid.MustParse(payload.WorkoutID),
		ExerciseID:  uuid.MustParse(payload.ExerciseID),
		Sets:        payload.Sets,
		Repetitions: payload.Repetitions,
		Duration:    *payload.Duration,
		Notes:       *payload.Notes,
	}
	saved, err := s.Repository.SaveWorkoutExercise(ctx, we)
	if err != nil {
		return nil, err
	}
	return toWorkoutExerciseResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *weService.GetPayload) (*weService.WorkoutExercise, error) {
	we, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &weService.NotFound{Message: "Workout exercise not found"}
		}
		return nil, err
	}
	return toWorkoutExerciseResponse(we), nil
}

func (s *Service) List(ctx context.Context, payload *weService.ListPayload) ([]*weService.WorkoutExercise, error) {
	wes, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*weService.WorkoutExercise
	for _, we := range wes {
		response = append(response, toWorkoutExerciseResponse(&we))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *weService.UpdateWorkoutExercisePayload) (*weService.WorkoutExercise, error) {
	we, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &weService.NotFound{Message: "Workout exercise not found"}
		}
		return nil, err
	}
	we.Sets = payload.Sets
	we.Repetitions = payload.Repetitions
	we.Duration = *payload.Duration
	we.Notes = *payload.Notes
	_, err = s.Repository.SaveWorkoutExercise(ctx, *we)
	if err != nil {
		return nil, err
	}
	return toWorkoutExerciseResponse(we), nil
}

func (s *Service) Delete(ctx context.Context, payload *weService.DeletePayload) error {
	we, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &weService.NotFound{Message: "Workout exercise not found"}
		}
		return err
	}
	return s.Repository.DeleteWorkoutExercise(ctx, we.ID.String())
}

func toWorkoutExerciseResponse(we *models.WorkoutExercise) *weService.WorkoutExercise {
	return &weService.WorkoutExercise{
		ID:          we.ID.String(),
		WorkoutID:   we.WorkoutID.String(),
		ExerciseID:  we.ExerciseID.String(),
		Sets:        we.Sets,
		Repetitions: we.Repetitions,
		Duration:    &we.Duration,
		Notes:       &we.Notes,
	}
}
