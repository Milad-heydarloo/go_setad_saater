package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var loginReq utils.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := fmt.Sprintf(`filter=organization_code="%s"`, loginReq.OrganizationCode)
	users, err := utils.MakeRequestToPocketBase(utils.PocketBaseURL+"?"+filter, "GET", nil)
	if err != nil || len(users) == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	user := users[0]
	storedPassword := user["password"].(string)
	if storedPassword != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

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
