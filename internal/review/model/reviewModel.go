package model

import (
	"github.com/google/uuid"
	userModel "github.com/pkstpm/Softdev-Backend/internal/users/model"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReservationID  uuid.UUID      `gorm:"type:uuid;not null;foreignkey:Reservation"`
	UserID         uuid.UUID      `gorm:"type:uuid;not null;foreignkey:User"`
	User           userModel.User `gorm:"foreignKey:UserID;references:ID"`
	Content        string         `gorm:"not null"`
	FoodRating     int            `gorm:"not null"`
	ServiceRating  int            `gorm:"not null"`
	AbbienceRating int            `gorm:"not null"`
}
