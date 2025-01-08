package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomCode تولید یک کد 5 رقمی تصادفی
func GenerateRandomCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(90000) + 10000
}

// تعریف ساختارهای مشترک
type User struct {
	OrganizationCode      string `json:"organization_code"`
	LandlineNumber        string `json:"landline_number"`
	Email                 string `json:"email"`
	Password              string `json:"password,omitempty"`
	FullName              string `json:"full_name"`
	OrganizationalAddress string `json:"organizational_address"`
	MobileNumber          string `json:"mobile_number"`
}

type LoginRequest struct {
	OrganizationCode string `json:"organization_code"`
	Password         string `json:"password"`
}

type Response struct {
	Message    string `json:"message"`
	RandomCode int    `json:"random_code,omitempty"`
	User       *User  `json:"user,omitempty"`
}
