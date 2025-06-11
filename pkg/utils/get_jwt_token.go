package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetJWTSecret() []byte {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY is not set")
	}
	return []byte(secret)
}
