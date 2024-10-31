package repository

import (
	"github.com/pkstpm/Softdev-Backend/internal/reservation/model"
)

type ReservationRepository interface {
	CreateReservation(reservation *model.Reservation) (*model.Reservation, error)
	CreateDishItem(dishItem *model.DishItem) (*model.DishItem, error)
	GetReservationById(reservationId string) (*model.Reservation, error)
	UpdateReservation(reservation *model.Reservation) error
	GetReservationByUserId(userId string) ([]model.Reservation, error)
	GetResvationByUserId(userId string) ([]model.Reservation, error)
	GetReservationByRestaurantId(restaurantId string) ([]model.Reservation, error)
	UpdateExpiredReservations() error
}
