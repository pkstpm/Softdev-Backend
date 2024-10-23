package service

import (
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
