package review

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null;foreignkey:Restaurant"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;foreignkey:User"`
	Content        string    `gorm:"not null"`
	FoodRating     int       `gorm:"not null"`
	ServiceRating  int       `gorm:"not null"`
	AbbienceRating int       `gorm:"not null"`
}
