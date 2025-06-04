package migration

import (
	service "pbl409-dashboard/pkg/services"

	"gorm.io/gorm"
)

func MigrateService(db *gorm.DB) error {
	err := db.AutoMigrate(&service.Service{})
	if err != nil {
		return err
	}
	return nil
}
