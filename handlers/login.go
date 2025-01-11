//

package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"net/http"
)

// LoginHandler handles login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	// بررسی متد درخواست
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// دیکود کردن بدنه درخواست
	var loginReq utils.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// بررسی وجود کاربر در PocketBase
	filter := fmt.Sprintf(`filter=organization_code="%s"`, loginReq.OrganizationCode)
	users, err := utils.MakeRequestToPocketBase(utils.PocketBaseURL+"?"+filter, "GET", nil)
	if err != nil || len(users) == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// بررسی رمز عبور
	user := users[0]
	storedPassword, ok := user["password"].(string)
	if !ok || storedPassword != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// ایجاد پاسخ موفقیت‌آمیز
	response := utils.Response{
		Message: "Login successful",
		User: &utils.User{
			OrganizationCode:      user["organization_code"].(string),
			LandlineNumber:        user["landline_number"].(string),
			Email:                 user["email"].(string),
			FullName:              user["full_name"].(string),
			OrganizationalAddress: user["organizational_address"].(string),
			MobileNumber:          user["mobile_number"].(string),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
