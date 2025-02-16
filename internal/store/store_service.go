package store

import (
	"aion/internal/database/db"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: db.DB.AionDB,
	}
}
