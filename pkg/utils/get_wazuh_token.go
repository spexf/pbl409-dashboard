package utils

import (
	"pbl409-dashboard/config"
	wazuh "pbl409-dashboard/pkg/wazuh/client"

	"github.com/go-redis/redis/v8"
)

func GetWazuhToken(host *wazuh.WazuhHost) (string, error) {
	token, err := config.RedisClient.Get(config.RedisContextBackground, "wazuh_token").Result()
	if err == redis.Nil {
		return wazuh.Authenticate(host, config.RedisClient)
	} else if err != nil {
		return "", err
	}
	return token, nil

}
