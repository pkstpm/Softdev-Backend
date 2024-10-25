package dto

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

type RegisterDTO struct {
	Username    string `json:"username" validate:"required,username,min=3,max=30"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,password"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
	DisplayName string `json:"displayName" validate:"required,min=3,max=30"`
}

type LoginDTO struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type RegisterRestaurantDTO struct {
	RestaurantName string `json:"restaurant_name" validate:"required,username,min=3,max=30"`
	RestaurantLoca string `json:"restaurant_location" validate:"required,username,min=3,max=30"`
	Category       string `json:"category" validate:"required,min=1,max=30"`
	Description    string `json:"description" validate:"required,min=3,max=30"`
}

type UserResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumebr string `json:"phone_number"`
	AccessToken string `json:"access_token"`
}

func CustomUsernameValidator(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	if len(username) < 3 || len(username) > 30 {
		return false
	}
	// Regular expression to ensure the username does not contain '@'
	re := regexp.MustCompile(`^[^@]+$`)
	return re.MatchString(username)
}

// CustomPasswordValidator is a custom password validation function
func CustomPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// Ensure password has at least one lowercase letter, one uppercase letter, one number, and one special character
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+\-=$begin:math:display$$end:math:display${};':"\\|,.<>/?]`).MatchString
	)
	if len(password) < 8 || len(password) > 30 {
		return false
	}
	if !hasUpper(password) || !hasLower(password) || !hasNumber(password) || !hasSpecial(password) {
		return false
	}
	return true
}

func CustomPhoneNumberValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	// Parse and validate phone number
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "TH")
	if err != nil {
		log.Printf("Invalid phone number format: %v", err)
		return false
	}

	// Check if the phone number is valid
	return phonenumbers.IsValidNumber(parsedNumber)
}

func RegisterValidators(validate *validator.Validate) {
	if err := validate.RegisterValidation("username", CustomUsernameValidator); err != nil {
		// Handle the error (e.g., log it, return it, etc.)
		log.Printf("Failed to register username validator: %v", err)
	}
	if err := validate.RegisterValidation("password", CustomPasswordValidator); err != nil {
		// Handle the error (e.g., log it, return it, etc.)
		log.Printf("Failed to register password validator: %v", err)
	}
	if err := validate.RegisterValidation("phoneNumber", CustomPhoneNumberValidator); err != nil {
		log.Printf("Failed to register phone validator: %v", err)
	}
}
