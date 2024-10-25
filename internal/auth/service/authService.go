package service

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/auth/dto"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
)

type AuthService interface {
	Register(dto *dto.RegisterDTO) error
	Login(dto *dto.LoginDTO) (string, error)
	RegisterRestaurant(userId uuid.UUID, dto *dto.RegisterRestaurantDTO) error
	Me(userId string) (*model.User, string, error)
}
