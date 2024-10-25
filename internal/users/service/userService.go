package service

import (
	"github.com/pkstpm/Softdev-Backend/internal/users/dto"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
)

type UserService interface {
	EditProfile(userId string, dto dto.EditProfileDTO) (*model.User, error)
	ChangePassword(userId string, dto dto.ChangePasswordDTO) error
	GetUser(userId string) (*model.User, error)
	UploadUserProfilePicture(userId string, url string) (*model.User, error)
}
