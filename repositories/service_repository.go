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

func GetService(db *gorm.DB) ([]models.Service, error) {
	var service []models.Service
	if err := db.Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func StoreService(db *gorm.DB, service *models.Service) error {
	if err := db.Create(service).Error; err != nil {
		return err
	}

	return nil
}

func DeleteService(db *gorm.DB, id uint) error {
	var service models.Service

	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&service, id).Error; err != nil {
		return err
	}
	return nil
}
