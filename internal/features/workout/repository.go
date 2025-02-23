package workout

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

func (r *Repository) FindByID(ctx context.Context, workoutID string) (*models.Workout, error) {
	var w models.Workout
	if err := r.DB.WithContext(ctx).Where("id = ?", workoutID).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &w, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.Workout, error) {
	var workouts []models.Workout
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (r *Repository) SaveWorkout(ctx context.Context, w models.Workout) (*models.Workout, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&w).Error; err != nil {
			utils.Log.Error(ctx, w, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *Repository) DeleteWorkout(ctx context.Context, workoutID string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		w, err := r.FindByID(ctx, workoutID)
		if err != nil {
			return err
		}
		if err := tx.Delete(w).Error; err != nil {
			return errors.New("error deleting workout")
		}
		return nil
	})
}
