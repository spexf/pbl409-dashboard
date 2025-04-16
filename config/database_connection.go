package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")
	dbSSL := os.Getenv("DATABASE_SSL")
	dbTimezone := os.Getenv("DATABASE_TIMEZONE")
	dbHost := os.Getenv("DATABASE_HOST")

	db, err := gorm.Open(postgres.Open("host=" + dbHost + " user= " + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSL + " TimeZone=" + dbTimezone))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
