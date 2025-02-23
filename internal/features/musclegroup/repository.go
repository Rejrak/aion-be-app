package musclegroup

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

func (r *Repository) FindByID(ctx context.Context, id string) (*models.MuscleGroup, error) {
	var mg models.MuscleGroup
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&mg).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &mg, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.MuscleGroup, error) {
	var groups []models.MuscleGroup
	if err := r.DB.WithContext(ctx).Limit(limit).Offset(offset).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *Repository) SaveMuscleGroup(ctx context.Context, mg models.MuscleGroup) (*models.MuscleGroup, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&mg).Error; err != nil {
			utils.Log.Error(ctx, mg, err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &mg, nil
}

func (r *Repository) DeleteMuscleGroup(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		mg, err := r.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if err := tx.Delete(mg).Error; err != nil {
			return errors.New("error deleting muscle group")
		}
		return nil
	})
}
