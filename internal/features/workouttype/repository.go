package workouttype

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

func (r *Repository) FindByID(ctx context.Context, id string) (*models.WorkoutType, error) {
	var wt models.WorkoutType
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&wt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &wt, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.WorkoutType, error) {
	var wts []models.WorkoutType
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&wts).Error; err != nil {
		return nil, err
	}
	return wts, nil
}

func (r *Repository) SaveWorkoutType(ctx context.Context, wt models.WorkoutType) (*models.WorkoutType, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&wt).Error; err != nil {
			utils.Log.Error(ctx, wt, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wt, nil
}

func (r *Repository) DeleteWorkoutType(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		wt, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(wt).Error; err != nil {
			return errors.New("error deleting workout type")
		}
		return nil
	})
}
