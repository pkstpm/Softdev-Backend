package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;foreignkey:User"`
	RestaurantID  uuid.UUID `gorm:"type:uuid;not null;foreignkey:Restaurant"`
	ReservationID uuid.UUID `gorm:"type:uuid;not null;foreignkey:Reservation"`
	Time          time.Time `gorm:"not null"`
	Content       string    `gorm:"not null"`
	IsRead        bool      `gorm:"not null"`
	Receiver      string    `gorm:"not null"`
}
