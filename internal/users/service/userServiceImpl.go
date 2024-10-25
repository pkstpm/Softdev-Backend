package service

import (
	"errors"

	"github.com/pkstpm/Softdev-Backend/internal/users/dto"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
	"github.com/pkstpm/Softdev-Backend/internal/users/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) EditProfile(userId string, editDTO dto.EditProfileDTO) (*model.User, error) {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	user.DisplayName = editDTO.DisplayName
	user.PhoneNumber = editDTO.PhoneNumber

	err = s.userRepository.EditProfileUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) GetUser(userId string) (*model.User, error) {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userServiceImpl) ChangePassword(userId string, dto dto.ChangePasswordDTO) error {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil || !user.CheckPassword(dto.Password) {
		return errors.New("invalid credentials")
	}

	user.Password, err = model.HashPassword(dto.NewPassword)

	if err != nil {
		return err
	}
	err = s.userRepository.EditProfileUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) UploadUserProfilePicture(userId string, url string) (*model.User, error) {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		return nil, err
	}
	user.ImgPath = url

	err = s.userRepository.EditProfileUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
