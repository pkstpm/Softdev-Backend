package service

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/dto"
)

type ReservationService interface {
	// GetReservationById(reservationId string) (*model.Reservation, error)
	// GetReservationsByUserId(userId string) ([]model.Reservation, error)
	// GetReservationsByRestaurantId(restaurantId string) ([]model.Reservation, error)
	CreateReservation(userId uuid.UUID, dto dto.CreateReservationDTO) error
}
