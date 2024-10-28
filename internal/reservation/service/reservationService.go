package service

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/dto"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/model"
)

type ReservationService interface {
	GetReservationById(reservationId string) (*model.Reservation, error)
	GetReservationsByUserId(userId string) ([]model.Reservation, error)
	// GetReservationsByRestaurantId(restaurantId string) ([]model.Reservation, error)
	CreateReservation(userId uuid.UUID, dto dto.CreateReservationDTO) (string, error)
	AddDishItem(userId string, reservationId string, dto dto.AddDishItemDTO) error
}
