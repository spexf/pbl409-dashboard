package service

import (
	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id uint) (*Service, error) {
	var service Service
	if err := db.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func Get(db *gorm.DB) ([]Service, error) {
	var service []Service
	if err := db.Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func Store(db *gorm.DB, service *Service) error {
	if err := db.Create(service).Error; err != nil {
		return err
	}

	return nil
}

func Delete(db *gorm.DB, id uint) error {
	var service Service

	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&service, id).Error; err != nil {
		return err
	}
	return nil
}

func Update(db *gorm.DB, id uint, updated map[string]interface{}) error {

	var service Service
	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Model(&service).Updates(updated).Error; err != nil {
		return err
	}

	return nil
}
