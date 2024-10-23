package service

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/dto"
)

type ReservationService interface {
	CreateReservation(userId uuid.UUID, dto dto.CreateReservationDTO) error
}
