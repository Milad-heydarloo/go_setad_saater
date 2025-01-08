package handlers

import (
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"net/http"
	"strconv"
)

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		MobileNumber string `json:"mobile_number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := fmt.Sprintf(`filter=mobile_number="%s"`, request.MobileNumber)
	users, err := utils.MakeRequestToPocketBase(utils.PocketBaseURL+"?"+filter, "GET", nil)
	if err != nil || len(users) == 0 {
		http.Error(w, "Mobile number not found. Please register first.", http.StatusNotFound)
		return
	}

	randomCode := utils.GenerateRandomCode()
	err = utils.SendVerificationCode("09190694410", "h826e7m", strconv.Itoa(randomCode), request.MobileNumber, 169397)
	if err != nil {
		http.Error(w, "Error sending verification code", http.StatusInternalServerError)
		return
	}

	response := utils.Response{
		Message:    "Verification code sent successfully",
		RandomCode: randomCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
