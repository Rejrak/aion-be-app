package workoutprogress

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

func (r *Repository) FindByID(ctx context.Context, id string) (*models.WorkoutProgress, error) {
	var wp models.WorkoutProgress
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&wp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &wp, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.WorkoutProgress, error) {
	var wps []models.WorkoutProgress
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&wps).Error; err != nil {
		return nil, err
	}
	return wps, nil
}

func (r *Repository) SaveWorkoutProgress(ctx context.Context, wp models.WorkoutProgress) (*models.WorkoutProgress, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&wp).Error; err != nil {
			utils.Log.Error(ctx, wp, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wp, nil
}

func (r *Repository) DeleteWorkoutProgress(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		wp, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(wp).Error; err != nil {
			return errors.New("error deleting workout progress")
		}
		return nil
	})
}
