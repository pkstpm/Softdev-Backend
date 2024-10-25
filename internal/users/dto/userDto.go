package dto

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

type EditProfileDTO struct {
	DisplayName string `json:"displayName" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
}

type ChangePasswordDTO struct {
	Password    string `json:"password" validate:"required,password"`
	NewPassword string `json:"newPassword" validate:"required,password"`
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

func RegisterValidators(validate *validator.Validate) {
	if err := validate.RegisterValidation("phoneNumber", CustomPhoneNumberValidator); err != nil {
		log.Printf("Failed to register phone validator: %v", err)
	}
	if err := validate.RegisterValidation("password", CustomPasswordValidator); err != nil {
		// Handle the error (e.g., log it, return it, etc.)
		log.Printf("Failed to register password validator: %v", err)
	}
}
