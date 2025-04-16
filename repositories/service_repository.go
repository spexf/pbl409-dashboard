package repositories

import (
	"pbl409-dashboard/models"

	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id uint) (*models.Service, error) {
	var service models.Service
	if err := db.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}
