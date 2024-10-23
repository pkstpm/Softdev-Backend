package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/auth/dto"
	restaurantModel "github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
	restaurantRepository "github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
	userModel "github.com/pkstpm/Softdev-Backend/internal/users/model"
	userRepository "github.com/pkstpm/Softdev-Backend/internal/users/repository"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

type authServiceImpl struct {
	userRepository       userRepository.UserRepository
	restaurantRepository restaurantRepository.RestaurantRepository
}

func NewAuthService(userRepository userRepository.UserRepository, restaurantRepository restaurantRepository.RestaurantRepository) AuthService {
	return &authServiceImpl{userRepository: userRepository, restaurantRepository: restaurantRepository}
}

func (s *authServiceImpl) Register(dto *dto.RegisterDTO) error {
	user := &userModel.User{
		Username:    dto.Username,
		Email:       dto.Email,
		Password:    dto.Password,
		PhoneNumber: dto.PhoneNumber,
		DisplayName: dto.DisplayName,
		UserType:    userModel.Customer,
	}

	err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) Login(dto *dto.LoginDTO) (string, error) {

	user, err := s.userRepository.FindUserByEmailOrUsername(dto.Identifier)
	if err != nil || !user.CheckPassword(dto.Password) {
		return "", errors.New("invalid credentials")
	}

	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("could not generate access token")
	}

	return accessToken, nil
}

func (s *authServiceImpl) RegisterRestaurant(userId uuid.UUID, dto *dto.RegisterRestaurantDTO) error {

	restaurant := &restaurantModel.Restaurant{
		UserID:         userId,
		RestaurantName: dto.RestaurantName,
		RestaurantLoca: dto.RestaurantLoca,
		Category:       dto.Category,
		Description:    dto.Description,
	}

	err := s.restaurantRepository.CreateRestaurant(restaurant)
	if err != nil {
		return err
	}

	err = s.userRepository.UpdateUserType(userId.String(), "Restaurant")
	if err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) Me(userId string) (string, error) {
	user, err := s.userRepository.FindUserByID(userId)
	if err != nil {
		return "", errors.New("user not found")
	}
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("could not generate access token")
	}

	return accessToken, nil
}
