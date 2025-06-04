package migration

import (
	user "pbl409-dashboard/pkg/users"

	"gorm.io/gorm"
)

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}
	return nil
}
