package migration

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Username string
	Password string
}

func MigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Database Migration Failed: %v", err)
	}
}
