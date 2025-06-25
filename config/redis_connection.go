package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var RedisContextBackground = context.Background()
var RedisClient *redis.Client

func InitRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // default value jika .env kosong
	}

	redisPassword := os.Getenv("REDIS_PASSWORD") // kalau ada password di .env

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err = RedisClient.Ping(RedisContextBackground).Result()
	if err != nil {
		log.Fatalf("Cannot connect to Redis: %v", err)
	}
}
