//

package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"io"
	"net/http"
)

// DeleteFileRequest ساختار برای دریافت ID فایل
type DeleteFileRequest struct {
	ID string `json:"id"`
}

// DeleteFileHandler هندلر برای حذف فایل با پشتیبانی از CORS
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	// تنظیم هدرهای CORS
	w.Header().Set("Access-Control-Allow-Origin", "https://setad.saaterco.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// پاسخ به درخواست OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// فقط اجازه درخواست POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"error":"Invalid request method"}`)
		fmt.Println("Invalid request method:", r.Method)
		return
	}

	// دیکود کردن بدنه درخواست
	var reqBody DeleteFileRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil || reqBody.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"Invalid or missing file ID"}`)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}

	// ساخت URL حذف فایل
	deleteURL := fmt.Sprintf("%s/%s", utils.PocketBaseFileURL, reqBody.ID)
	fmt.Printf("Attempting to delete file with ID: %s\n", reqBody.ID)
	fmt.Printf("DELETE URL: %s\n", deleteURL)

	// ارسال درخواست DELETE به PocketBase
	req, err := http.NewRequest(http.MethodDelete, deleteURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Failed to create delete request: %s"}`, err)
		fmt.Printf("Error creating DELETE request: %v\n", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Failed to send delete request: %s"}`, err)
		fmt.Printf("Error sending DELETE request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// بررسی پاسخ
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Failed to delete file: %s"}`, body)
		fmt.Printf("Failed to delete file. Status: %d, Response: %s\n", resp.StatusCode, string(body))
		return
	}

	// موفقیت‌آمیز
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"File deleted successfully"}`)
	fmt.Printf("File with ID %s deleted successfully\n", reqBody.ID)
}
