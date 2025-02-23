package workoutexercise

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

func (r *Repository) FindByID(ctx context.Context, id string) (*models.WorkoutExercise, error) {
	var we models.WorkoutExercise
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&we).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &we, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.WorkoutExercise, error) {
	var wes []models.WorkoutExercise
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&wes).Error; err != nil {
		return nil, err
	}
	return wes, nil
}

func (r *Repository) SaveWorkoutExercise(ctx context.Context, we models.WorkoutExercise) (*models.WorkoutExercise, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&we).Error; err != nil {
			utils.Log.Error(ctx, we, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &we, nil
}

func (r *Repository) DeleteWorkoutExercise(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		we, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(we).Error; err != nil {
			return errors.New("error deleting workout exercise")
		}
		return nil
	})
}
