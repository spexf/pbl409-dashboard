package services

import (
	"pbl409-dashboard/models"
	"pbl409-dashboard/repositories"

	"gorm.io/gorm"
)

func GetService(db *gorm.DB) ([]models.Service, error) {
	return repositories.GetService(db)
}

func ShowService(db *gorm.DB, id uint) (*models.Service, error) {
	return repositories.FindById(db, id)
}
