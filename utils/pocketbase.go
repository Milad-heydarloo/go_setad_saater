package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const PocketBaseURL = "http://80.210.37.118:9000/api/collections/user_setad/records"
const PocketBaseFileURL = "http://80.210.37.118:9000/api/collections/file/records"
const PocketBaseBaseURL = "http://80.210.37.118:9000/api/collections/order_setad/records"

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

func CreateOrder(data map[string]interface{}) (string, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", PocketBaseBaseURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var result struct {
		Id string `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Id, nil
}

// UpdateOrderFiles updates the 'file' field of a specific order
func UpdateOrderFiles(orderID string, fileIDs []string) error {
	url := fmt.Sprintf("%s/%s", PocketBaseBaseURL, orderID) // ایجاد URL برای درخواست
	data := map[string]interface{}{
		"file": fileIDs,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var response struct {
		ID string `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	if response.ID == "" {
		return fmt.Errorf("update failed: no ID returned")
	}

	return nil
}

// UpdateOrderDescription updates the 'description' field of a specific order
func UpdateOrderDescription(orderID string, description string) error {
	url := fmt.Sprintf("%s/%s", PocketBaseBaseURL, orderID) // ساخت URL برای آپدیت
	data := map[string]interface{}{
		"description": description,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var response struct {
		ID string `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	if response.ID == "" {
		return fmt.Errorf("update failed: no ID returned")
	}

	return nil
}

// DeleteOrder removes an order by its ID from PocketBase
func DeleteOrder(orderID string) error {
	url := fmt.Sprintf("%s/%s", PocketBaseBaseURL, orderID) // ساخت URL حذف
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete order: HTTP error %d", resp.StatusCode)
	}

	return nil
}

// GetOrder retrieves a single order by its ID with expanded file information//
func GetOrder(orderID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s?expand=file", PocketBaseBaseURL, orderID) // اکسپند فایل‌ها در درخواست
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch order: HTTP error %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetOrdersByUser retrieves all orders associated with a specific user ID
func GetOrdersByUser(userID string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s?filter=user='%s'&expand=file,invoice_file,payment_receipt_file", PocketBaseBaseURL, userID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch orders: HTTP error %d", resp.StatusCode)
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
