package model

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;foreignkey:User"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null;foreignkey:Restaurant"`
	Content      string    `gorm:"not null"`
	IsRead       bool      `gorm:"not null"`
	Sender       string    `gorm:"not null"`
}
