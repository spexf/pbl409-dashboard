package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisContextBackground = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "cache:6379",
	Password: "",
	DB:       0,
})

func InitRedis() {

	_, err := RedisClient.Ping(RedisContextBackground).Result()
	if err != nil {
		log.Fatalf("Cannot connect to Redis")
	}
}
