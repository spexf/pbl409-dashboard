package wazuh

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type WazuhHost struct {
	ServiceID uint
	Username  string
	Password  string
	Host      string
}

type AuthResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Error int `json:"error"`
}

var WazuhBaseURL = "https://%s:8081%s"

func Authenticate(cred *WazuhHost, rdb *redis.Client) (string, error) {
	url := fmt.Sprintf(WazuhBaseURL, cred.Host, "/security/user/authenticate")

	auth := cred.Username + ":" + cred.Password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", basicAuth)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth failed: %s", string(body))
	}

	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", err
	}

	token := authResp.Data.Token
	// Simpan token ke Redis selama 15 menit
	err = rdb.Set(context.Background(), "wazuh_token", token, 15*time.Minute).Err()
	if err != nil {
		return "", fmt.Errorf("failed to store token in redis: %v", err)
	}

	return token, nil
}
