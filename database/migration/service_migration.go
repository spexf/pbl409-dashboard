package migration

import (
	"pbl409-dashboard/models"

	"gorm.io/gorm"
)

func MigrateService(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Service{})
	if err != nil {
		return err
	}
	return nil
}
