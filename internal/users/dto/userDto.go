package dto

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

type EditProfileDTO struct {
	DisplayName string `json:"displayName" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber"`
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
	if err := validate.RegisterValidation("phoneNumber", CustomPhoneNumberValidator); err != nil {
		log.Printf("Failed to register phone validator: %v", err)
	}
}
