package services

import (
	"pbl409-dashboard/dtos"
	"pbl409-dashboard/models"
	"pbl409-dashboard/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetService(db *gorm.DB) ([]models.Service, error) {
	return repositories.GetService(db)
}

func ShowService(db *gorm.DB, id uint) (*models.Service, error) {
	return repositories.FindById(db, id)
}

func StoreService(db *gorm.DB, store dtos.ServiceStore) error {
	if err := validate.Struct(store); err != nil {
		return err
	}

	service := models.Service{
		Name:     store.Name,
		Host:     store.Host,
		Type:     store.Type,
		Username: store.Username,
		Password: store.Password,
	}

	return repositories.StoreService(db, &service)
}
