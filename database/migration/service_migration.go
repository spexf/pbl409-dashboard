package migration

import (
	"log"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string
	Host        string
	Type        string
	Username    string
	Password    string
	ActiveToken string
}

func MigrateService(db *gorm.DB) {
	err := db.AutoMigrate(&Service{})
	if err != nil {
		log.Fatalf("Database Migration Failed: %v", err)
	}
}
