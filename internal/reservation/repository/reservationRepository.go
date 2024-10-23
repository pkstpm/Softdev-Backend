package repository

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/model"
)

type ReservationRepository interface {
	CreateReservation(reservation *model.Reservation) (uuid.UUID, error)
	CreateDishItem(dishItem *model.DishItem) error
}
