package trainingplan

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"
	"time"

	trainingPlanService "be/gen/trainingplan"

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

// Create crea un nuovo TrainingPlan
func (s *Service) Create(ctx context.Context, payload *trainingPlanService.CreateTrainingPlanPayload) (*trainingPlanService.TrainingPlan, error) {
	plan := models.TrainingPlan{
		Name:          payload.Name,
		Description:   *payload.Description,
		StartDate:     parseTime(payload.StartDate),
		EndDate:       parseTime(payload.EndDate),
		UserID:        uuid.MustParse(payload.UserID),
		WorkoutTypeID: uuid.MustParse(payload.WorkoutTypeID),
	}

	saved, err := s.Repository.SaveTrainingPlan(ctx, plan)
	if err != nil {
		return nil, err
	}

	return toTrainingPlanResponse(saved), nil
}

// Get restituisce un TrainingPlan per ID
func (s *Service) Get(ctx context.Context, payload *trainingPlanService.GetPayload) (*trainingPlanService.TrainingPlan, error) {
	plan, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &trainingPlanService.NotFound{Message: "Training plan not found"}
		}
		return nil, err
	}
	return toTrainingPlanResponse(plan), nil
}

// List restituisce un elenco di TrainingPlan con paginazione
func (s *Service) List(ctx context.Context, payload *trainingPlanService.ListPayload) ([]*trainingPlanService.TrainingPlan, error) {
	plans, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*trainingPlanService.TrainingPlan
	for _, plan := range plans {
		response = append(response, toTrainingPlanResponse(&plan))
	}
	return response, nil
}

// Update aggiorna un TrainingPlan esistente
func (s *Service) Update(ctx context.Context, payload *trainingPlanService.UpdateTrainingPlanPayload) (*trainingPlanService.TrainingPlan, error) {
	plan, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &trainingPlanService.NotFound{Message: "Training plan not found"}
		}
		return nil, err
	}

	plan.Name = payload.Name
	plan.Description = *payload.Description
	plan.StartDate = parseTime(payload.StartDate)
	plan.EndDate = parseTime(payload.EndDate)
	plan.UserID = uuid.MustParse(payload.UserID)
	plan.WorkoutTypeID = uuid.MustParse(payload.WorkoutTypeID)

	_, err = s.Repository.SaveTrainingPlan(ctx, *plan)
	if err != nil {
		return nil, err
	}
	return toTrainingPlanResponse(plan), nil
}

// Delete elimina un TrainingPlan
func (s *Service) Delete(ctx context.Context, payload *trainingPlanService.DeletePayload) error {
	plan, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &trainingPlanService.NotFound{Message: "Training plan not found"}
		}
		return err
	}
	return s.Repository.DeleteTrainingPlan(ctx, plan.ID.String())
}

// helper: conversione modello -> response DSL
func toTrainingPlanResponse(plan *models.TrainingPlan) *trainingPlanService.TrainingPlan {
	return &trainingPlanService.TrainingPlan{
		ID:            plan.ID.String(),
		Name:          plan.Name,
		Description:   &plan.Description,
		StartDate:     plan.StartDate.Format(time.RFC3339),
		EndDate:       plan.EndDate.Format(time.RFC3339),
		UserID:        plan.UserID.String(),
		WorkoutTypeID: plan.WorkoutTypeID.String(),
	}
}

func parseTime(ts string) time.Time {
	t, _ := time.Parse(time.RFC3339, ts)
	return t
}
