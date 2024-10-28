package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateReservationDTO struct {
	RestaurantID uuid.UUID `json:"restaurant_id" validate:"required"`
	TableID      uuid.UUID `json:"table_id" validate:"required"`
	StartTime    time.Time `json:"start_time" validate:"required"`
	EndTime      time.Time `json:"end_time" validate:"required"`
}

type AddDishItemDTO struct {
	DishItems []DishItemDTO `json:"dish_items" validate:"required,dive"`
}

type DishItemDTO struct {
	DishID   uuid.UUID `json:"dish_id" validate:"required"`
	Quantity int       `json:"quantity" validate:"required"`
	Option   string    `json:"option"`
	Comment  string    `json:"comment"`
}
