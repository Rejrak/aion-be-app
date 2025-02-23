package exercise

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

func (r *Repository) FindByID(ctx context.Context, id string) (*models.Exercise, error) {
	var ex models.Exercise
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&ex).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &ex, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.Exercise, error) {
	var exs []models.Exercise
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&exs).Error; err != nil {
		return nil, err
	}
	return exs, nil
}

func (r *Repository) SaveExercise(ctx context.Context, ex models.Exercise) (*models.Exercise, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&ex).Error; err != nil {
			utils.Log.Error(ctx, ex, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

func (r *Repository) DeleteExercise(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ex, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(ex).Error; err != nil {
			return errors.New("error deleting exercise")
		}
		return nil
	})
}
