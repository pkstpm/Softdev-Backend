package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserType string

const (
	Customer   UserType = "Customer"
	Restaurant UserType = "Restaurant"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username    string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	PhoneNumber string    `gorm:"not null"`
	DisplayName string    `gorm:"not null"`
	UserType    UserType  `gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	if len(u.Password) > 0 {
		u.Password, err = hashPassword(u.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (u *User) CheckPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil // return true if the password matches
}
