package workoutprogress

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"
	"time"

	wpService "be/gen/workoutprogress"

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

func (s *Service) Create(ctx context.Context, payload *wpService.CreateWorkoutProgressPayload) (*wpService.WorkoutProgress, error) {
	wp := models.WorkoutProgress{
		WorkoutID: uuid.MustParse(payload.WorkoutID),
		UserID:    uuid.MustParse(payload.UserID),
		Date:      parseTime(payload.Date),
	}
	saved, err := s.Repository.SaveWorkoutProgress(ctx, wp)
	if err != nil {
		return nil, err
	}
	return toWorkoutProgressResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *wpService.GetPayload) (*wpService.WorkoutProgress, error) {
	wp, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &wpService.NotFound{Message: "Workout progress not found"}
		}
		return nil, err
	}
	return toWorkoutProgressResponse(wp), nil
}

func (s *Service) List(ctx context.Context, payload *wpService.ListPayload) ([]*wpService.WorkoutProgress, error) {
	wps, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*wpService.WorkoutProgress
	for _, wp := range wps {
		response = append(response, toWorkoutProgressResponse(&wp))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *wpService.UpdateWorkoutProgressPayload) (*wpService.WorkoutProgress, error) {
	wp, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &wpService.NotFound{Message: "Workout progress not found"}
		}
		return nil, err
	}
	wp.Date = parseTime(payload.Date)
	_, err = s.Repository.SaveWorkoutProgress(ctx, *wp)
	if err != nil {
		return nil, err
	}
	return toWorkoutProgressResponse(wp), nil
}

func (s *Service) Delete(ctx context.Context, payload *wpService.DeletePayload) error {
	wp, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &wpService.NotFound{Message: "Workout progress not found"}
		}
		return err
	}
	return s.Repository.DeleteWorkoutProgress(ctx, wp.ID.String())
}

func toWorkoutProgressResponse(wp *models.WorkoutProgress) *wpService.WorkoutProgress {
	return &wpService.WorkoutProgress{
		ID:        wp.ID.String(),
		WorkoutID: wp.WorkoutID.String(),
		UserID:    wp.UserID.String(),
		Date:      wp.Date.Format(time.RFC3339),
	}
}

func parseTime(ts string) time.Time {
	t, _ := time.Parse(time.RFC3339, ts)
	return t
}
