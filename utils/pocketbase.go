package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const PocketBaseURL = "http://80.210.37.118:9000/api/collections/user_setad/records"

func MakeRequestToPocketBase(url, method string, body []byte) ([]map[string]interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var result struct {
		Items []map[string]interface{} `json:"items"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
