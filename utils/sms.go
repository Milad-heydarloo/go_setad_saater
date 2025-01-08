package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const apiURL = "https://rest.payamak-panel.com/api/SendSMS/"

func SendVerificationCode(username, password, text, to string, bodyId int64) error {
	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(apiURL+"BaseServiceNumber", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return fmt.Errorf("failed to send SMS: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %d", response.StatusCode)
	}

	return nil
}
