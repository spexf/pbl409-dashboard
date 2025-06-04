package service

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetService(db *gorm.DB) ([]Service, error) {
	return Get(db)
}

func ShowService(db *gorm.DB, id uint) (*Service, error) {
	return FindById(db, id)
}

func DeleteService(db *gorm.DB, id uint) error {
	return Delete(db, id)
}

func UpdateService(db *gorm.DB, id uint, updated map[string]interface{}) error {
	return Update(db, id, updated)
}

func StoreService(db *gorm.DB, store ServiceStore) error {
	if err := validate.Struct(store); err != nil {
		return err
	}

	service := Service{
		Name:     store.Name,
		Host:     store.Host,
		Type:     store.Type,
		Username: store.Username,
		Password: store.Password,
	}

	return Store(db, &service)
}
