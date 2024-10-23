package model

import (
	"github.com/google/uuid"
	reservationModel "github.com/pkstpm/Softdev-Backend/internal/reservation/model"
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;foreignkey:User;unique"`
	RestaurantName string    `gorm:"not null"`
	RestaurantLoca string    `gorm:"not null"`
	Category       string    `gorm:"not null"`
	Description    string
	ImagePath      string
	Tables         []Table                        `gorm:"foreignKey:RestaurantID"`
	Reservations   []reservationModel.Reservation `gorm:"foreignKey:RestaurantID"`
	Dishes         []Dish                         `gorm:"foreignKey:RestaurantID"`
	TimeSlots      []TimeSlot                     `gorm:"foreignKey:RestaurantID"`
	// Reviews        []Review      `gorm:"foreignKey:RestaurantID"`
}

type Table struct {
	ID           uuid.UUID                      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID                      `gorm:"type:uuid;not null;foreignkey:Restaurant;"`
	TableNumber  int                            `gorm:"not null"`
	Capacity     int                            `gorm:"not null"`
	Reservations []reservationModel.Reservation `gorm:"foreignKey:TableID" json:"-"` // One-to-many relationship with Reservation
}

type TimeSlot struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Weekday      int        `gorm:"not null"`                         // Stores weekday as an integer (e.g., 0 for Sunday, 1 for Monday)
	HourStart    int        `gorm:"not null"`                         // Stores the starting hour as a DateTime
	HourEnd      int        `gorm:"not null"`                         // Stores the ending hour as a DateTime
	RestaurantID uuid.UUID  `gorm:"not null"`                         // Foreign key to Restaurant
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID" json:"-"` // Relation to Restaurant model
}

type Dish struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null;foreignkey:Restaurant;"` // Foreign key for Restaurant
	Name         string    `gorm:"not null"`                                  // Name of the dish
	Price        int       `gorm:"not null"`                                  // Price of the dish
	Description  string    `gorm:"not null"`                                  // Description of the dish
	OptionList   string    `gorm:"not null"`                                  // Options for the dish, stored as a string
}
