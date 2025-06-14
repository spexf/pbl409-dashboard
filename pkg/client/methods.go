package wazuh

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func WazuhGet(host *WazuhHost, token string, endpoint string) ([]byte, error) {
	url := fmt.Sprintf(WazuhBaseURL, host.Host, endpoint)

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", string(body))
	}

	return body, nil
}
