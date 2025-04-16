package migration

import (
	"pbl409-dashboard/models"

	"gorm.io/gorm"
)

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}
