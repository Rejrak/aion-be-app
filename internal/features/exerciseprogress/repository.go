package exerciseprogress

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
	return &Repository{DB: db}
}

func (r *Repository) FindByID(ctx context.Context, id string) (*models.ExerciseProgress, error) {
	var ep models.ExerciseProgress
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&ep).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &ep, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.ExerciseProgress, error) {
	var eps []models.ExerciseProgress
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&eps).Error; err != nil {
		return nil, err
	}
	return eps, nil
}

func (r *Repository) SaveExerciseProgress(ctx context.Context, ep models.ExerciseProgress) (*models.ExerciseProgress, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&ep).Error; err != nil {
			utils.Log.Error(ctx, ep, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ep, nil
}

func (r *Repository) DeleteExerciseProgress(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ep, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(ep).Error; err != nil {
			return errors.New("error deleting exercise progress")
		}
		return nil
	})
}
