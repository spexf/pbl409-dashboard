package services

import (
	"pbl409-dashboard/models"
	"pbl409-dashboard/repositories"

	"gorm.io/gorm"
)

func ShowService(db *gorm.DB, id uint) (*models.Service, error) {
	return repositories.FindById(db, id)
}
