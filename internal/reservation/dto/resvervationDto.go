package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateReservationDTO struct {
	RestaurantID uuid.UUID           `json:"restaurant_id" validate:"required"`
	TableID      uuid.UUID           `json:"table_id" validate:"required"`
	StartTime    time.Time           `json:"start_time" validate:"required"`
	EndTime      time.Time           `json:"end_time" validate:"required"`
	DishItems    []CreateDishItemDTO `json:"dish_items,omitempty"`
	TotalPrice   int                 `json:"total_price" validate:"required"`
}

type CreateDishItemDTO struct {
	DishID   uuid.UUID `json:"dish_id"`
	Quantity int       `json:"quantity"`
	Option   string    `json:"option"`
	Comment  string    `json:"comment"`
}
