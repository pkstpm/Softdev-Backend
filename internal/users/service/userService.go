package service

import (
	"github.com/pkstpm/Softdev-Backend/internal/users/dto"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
)

type UserService interface {
	EditProfile(userId string, dto dto.EditProfileDTO) (*model.User, error)
	GetUser(userId string) (*model.User, error)
}
