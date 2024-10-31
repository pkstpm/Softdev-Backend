package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
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

func (s *userServiceImpl) GetProfile(userId string) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Role:        string(user.UserType),
		PhoneNumber: user.PhoneNumber,
		ImgPath:     user.ImgPath,
	}, nil
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

func (s *userServiceImpl) AddFavouriteRestaurant(userId string, restaurantId string) error {
	log.Println(userId)
	log.Println(restaurantId)
	if s.userRepository.SearchFavouriteRestaurant(userId, restaurantId) {
		return errors.New("restaurant already in favourite")
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %w", err)
	}

	parsedRestaurantId, err := uuid.Parse(restaurantId)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %w", err)
	}

	favourite := &model.Favourite{
		UserID:       parsedUserId,
		RestaurantID: parsedRestaurantId,
	}

	err = s.userRepository.AddFavouriteRestaurant(favourite)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) RemoveFavouriteRestaurant(userId string, restaurantId string) error {
	if !s.userRepository.SearchFavouriteRestaurant(userId, restaurantId) {
		return errors.New("restaurant not in favourite")
	}

	err := s.userRepository.RemoveFavouriteRestaurant(restaurantId)
	if err != nil {
		return err
	}

	return nil
}
