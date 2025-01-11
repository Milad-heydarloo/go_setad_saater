//

package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"net/http"
)

// UpdateUserHandler handles updating user information
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// تنظیم هدرهای CORS
	enableCORSForUpdateUser(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// بررسی متد HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// ساختار درخواست
	var updateRequest struct {
		ID                    string `json:"id"`
		OrganizationalAddress string `json:"organizational_address,omitempty"`
		LandlineNumber        string `json:"landline_number,omitempty"`
	}

	// دیکود کردن بدنه درخواست
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil || updateRequest.ID == "" {
		http.Error(w, "Invalid request body or missing user ID", http.StatusBadRequest)
		return
	}

	// آماده‌سازی داده‌های به‌روزرسانی
	data := map[string]interface{}{}
	if updateRequest.OrganizationalAddress != "" {
		data["organizational_address"] = updateRequest.OrganizationalAddress
	}
	if updateRequest.LandlineNumber != "" {
		data["landline_number"] = updateRequest.LandlineNumber
	}

	// اگر هیچ داده‌ای برای به‌روزرسانی ارسال نشده باشد
	if len(data) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	// ارسال درخواست به PocketBase
	url := fmt.Sprintf("%s/%s", utils.PocketBaseURL, updateRequest.ID)
	body, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating update payload", http.StatusInternalServerError)
		return
	}

	_, err = utils.MakeRequestToPocketBase(url, "PATCH", body)
	if err != nil {
		http.Error(w, "Error updating user in PocketBase", http.StatusInternalServerError)
		return
	}

	// پاسخ موفقیت‌آمیز
	response := utils.Response{
		Message: "User updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// enableCORSForUpdateUser adds the CORS headers specific to UpdateUserHandler
func enableCORSForUpdateUser(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://setad.saaterco.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
