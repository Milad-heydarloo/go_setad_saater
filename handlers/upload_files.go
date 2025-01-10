package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"time"

	"go_setad_saater/utils"
)

// FileUploadResponse ساختار برای پاسخ شامل ID فایل‌ها
type FileUploadResponse struct {
	FileIDs []string `json:"file_ids"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	// تنظیم هدرهای CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Received a new upload request")
	fmt.Printf("Request Method: %s\n", r.Method)
	fmt.Printf("Request Content-Type: %s\n", r.Header.Get("Content-Type"))

	// بررسی و پردازش بدنه درخواست به عنوان multipart/form-data
	err := r.ParseMultipartForm(10 << 20) // محدودیت 10 مگابایت
	if err != nil {
		fmt.Printf("Error parsing form: %v\n", err)
		writeErrorResponse(w, http.StatusBadRequest, "Invalid multipart form data")
		return
	}

	var fileIDs []string
	for key, files := range r.MultipartForm.File {
		fmt.Printf("Field: %s\n", key)
		for _, fileHeader := range files {
			fmt.Printf("Processing file: %s\n", fileHeader.Filename)

			fileID, err := uploadSingleFileFromHeader(fileHeader)
			if err != nil {
				fmt.Printf("Error uploading file %s: %v\n", fileHeader.Filename, err)
				writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to upload file %s: %v", fileHeader.Filename, err))
				return
			}
			fileIDs = append(fileIDs, fileID)
		}
	}

	response := FileUploadResponse{FileIDs: fileIDs}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Printf("Successfully uploaded files: %v\n", fileIDs)
}

func uploadSingleFileFromHeader(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

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

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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
