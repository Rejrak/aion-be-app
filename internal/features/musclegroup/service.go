package musclegroup

import (
	"be/internal/database/db"
	"be/internal/database/models"
	"context"
	"errors"

	muscleGroupService "be/gen/musclegroup"

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

func (s *Service) Create(ctx context.Context, payload *muscleGroupService.CreateMuscleGroupPayload) (*muscleGroupService.MuscleGroup, error) {
	mg := models.MuscleGroup{
		Name:        payload.Name,
		Description: *payload.Description,
	}
	saved, err := s.Repository.SaveMuscleGroup(ctx, mg)
	if err != nil {
		return nil, err
	}
	return toMuscleGroupResponse(saved), nil
}

func (s *Service) Get(ctx context.Context, payload *muscleGroupService.GetPayload) (*muscleGroupService.MuscleGroup, error) {
	mg, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &muscleGroupService.NotFound{Message: "Muscle group not found"}
		}
		return nil, err
	}
	return toMuscleGroupResponse(mg), nil
}

func (s *Service) List(ctx context.Context, payload *muscleGroupService.ListPayload) ([]*muscleGroupService.MuscleGroup, error) {
	groups, err := s.Repository.List(ctx, payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}
	var response []*muscleGroupService.MuscleGroup
	for _, mg := range groups {
		response = append(response, toMuscleGroupResponse(&mg))
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, payload *muscleGroupService.UpdateMuscleGroupPayload) (*muscleGroupService.MuscleGroup, error) {
	mg, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &muscleGroupService.NotFound{Message: "Muscle group not found"}
		}
		return nil, err
	}
	mg.Name = payload.Name
	mg.Description = *payload.Description
	_, err = s.Repository.SaveMuscleGroup(ctx, *mg)
	if err != nil {
		return nil, err
	}
	return toMuscleGroupResponse(mg), nil
}

func (s *Service) Delete(ctx context.Context, payload *muscleGroupService.DeletePayload) error {
	mg, err := s.Repository.FindByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &muscleGroupService.NotFound{Message: "Muscle group not found"}
		}
		return err
	}
	return s.Repository.DeleteMuscleGroup(ctx, mg.ID.String())
}

func toMuscleGroupResponse(mg *models.MuscleGroup) *muscleGroupService.MuscleGroup {
	return &muscleGroupService.MuscleGroup{
		ID:          mg.ID.String(),
		Name:        mg.Name,
		Description: &mg.Description,
	}
}
