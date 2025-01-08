package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils" // مسیر صحیح به بسته utils
	"net/http"
	"strconv"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user utils.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := fmt.Sprintf(`filter=organization_code="%s"`, user.OrganizationCode)
	users, err := utils.MakeRequestToPocketBase(utils.PocketBaseURL+"?"+filter, "GET", nil)
	if err == nil && len(users) > 0 {
		response := utils.Response{
			Message: "User already exists. Please log in.",
			User:    &user,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	userData, _ := json.Marshal(user)
	_, err = utils.MakeRequestToPocketBase(utils.PocketBaseURL, "POST", userData)
	if err != nil {
		http.Error(w, "Error saving user to PocketBase", http.StatusInternalServerError)
		return
	}

	randomCode := utils.GenerateRandomCode()
	err = utils.SendVerificationCode("09190694410", "h826e7m", strconv.Itoa(randomCode), user.MobileNumber, 169397)
	if err != nil {
		http.Error(w, "Error sending verification code", http.StatusInternalServerError)
		return
	}

	response := utils.Response{
		Message:    "User registered successfully",
		RandomCode: randomCode,
		User:       &user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
