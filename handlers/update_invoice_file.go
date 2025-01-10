package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

// UpdateInvoiceFileHandler handles uploading a file and updating the invoice_file field for an order
func UpdateInvoiceFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// تنظیم هدرهای CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// حذف محدودیت حجم فایل
	if err := r.ParseMultipartForm(0); err != nil { // 0 یعنی بدون محدودیت
		http.Error(w, "Invalid multipart form data", http.StatusBadRequest)
		return
	}

	// Read order ID
	orderID := r.FormValue("order_id")
	if orderID == "" {
		http.Error(w, "Missing order_id", http.StatusBadRequest)
		return
	}

	// Read file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Missing file in request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload file to PocketBase
	fileID, err := uploadInvoiceFile(file, fileHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload file: %v", err), http.StatusInternalServerError)
		return
	}

	// Update the invoice_file field in the order
	err = updateInvoiceFile(orderID, fileID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update order: %v", err), http.StatusInternalServerError)
		return
	}

	// Success response
	response := map[string]string{
		"message":  "Invoice file updated successfully",
		"order_id": orderID,
		"file_id":  fileID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// uploadInvoiceFile uploads a file to PocketBase and returns the file ID
func uploadInvoiceFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("field", fileHeader.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}
	writer.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", utils.PocketBaseFileURL, body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("PocketBase returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var response struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return response.ID, nil
}

// updateInvoiceFile updates the invoice_file field for an order in PocketBase
func updateInvoiceFile(orderID, fileID string) error {
	url := fmt.Sprintf("%s/%s", utils.PocketBaseBaseURL, orderID)
	data := map[string]interface{}{
		"invoice_file": fileID,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal update data: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("PocketBase returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
