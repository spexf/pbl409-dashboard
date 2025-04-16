package migration

import (
	"log"

	"pbl409-dashboard/config"
)

func MigrateAll() {
	var error []error
	db := config.ConnectDB()
	if err := MigrateService(db); err != nil {
		error = append(error, err)
	}
	if err := MigrateUser(db); err != nil {
		error = append(error, err)
	}

	if error != nil {
		log.Fatalf("Migration Failed: %v", error)
	}

	log.Println("Migration Success !!")
}
