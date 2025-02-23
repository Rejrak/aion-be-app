package trainingplan

import (
	"be/internal/database/models"
	"be/internal/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindByID(ctx context.Context, planID string) (*models.TrainingPlan, error) {
	var plan models.TrainingPlan
	if err := r.DB.WithContext(ctx).Where("id = ?", planID).First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &plan, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.TrainingPlan, error) {
	var plans []models.TrainingPlan
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *Repository) SaveTrainingPlan(ctx context.Context, plan models.TrainingPlan) (*models.TrainingPlan, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&plan).Error; err != nil {
			utils.Log.Error(ctx, plan, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *Repository) DeleteTrainingPlan(ctx context.Context, planID string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		plan, err := r.FindByID(ctx, planID)
		if err != nil {
			return err
		}
		if err := tx.Delete(plan).Error; err != nil {
			return errors.New("error deleting training plan")
		}
		return nil
	})
}
