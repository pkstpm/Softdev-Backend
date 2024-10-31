package service

import (
	"log"

	"github.com/pkstpm/Softdev-Backend/internal/notification/dto"
	"github.com/pkstpm/Softdev-Backend/internal/notification/repository"
	restaurantRepository "github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
)

type notificationServiceImpl struct {
	notificationRepository repository.NotificationRepository
	restaurantRepository   restaurantRepository.RestaurantRepository
}

func NewNotificationService(notificationRepository repository.NotificationRepository, restaurantRepository restaurantRepository.RestaurantRepository) NotificationService {
	return &notificationServiceImpl{notificationRepository: notificationRepository, restaurantRepository: restaurantRepository}
}

func (s *notificationServiceImpl) GetUserNotReadNotification(userId string) ([]dto.NotificationResponse, error) {
	notifications, err := s.notificationRepository.GetUserNotReadNotification(userId)
	if err != nil {
		return nil, err
	}

	// Convert model notifications to NotificationResponse DTOs
	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			NotificationId: notification.ID.String(),            // Assuming ID is of type uuid.UUID
			ReservationId:  notification.ReservationID.String(), // Change this if it's not a UUID
			Content:        notification.Content,
			From:           notification.FromName,
			Time:           notification.Time,
		})
		notification.IsRead = true
		s.notificationRepository.UpdateNotification(&notification)
	}

	return notificationResponses, nil
}

func (s *notificationServiceImpl) GetRestaurantNotReadNotification(userId string) ([]dto.NotificationResponse, error) {
	restaurant, err := s.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil || restaurant == nil {
		return nil, err
	}
	notifications, err := s.notificationRepository.GetRestaurantNotReadNotification(restaurant.ID.String())
	if err != nil {
		return nil, err
	}
	log.Printf("restaurant: %v", restaurant)

	// Convert model notifications to NotificationResponse DTOs
	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			NotificationId: notification.ID.String(),            // Assuming ID is of type uuid.UUID
			ReservationId:  notification.ReservationID.String(), // Change this if it's not a UUID
			Content:        notification.Content,
			From:           notification.FromName,
			Time:           notification.Time,
		})
		notification.IsRead = true
		s.notificationRepository.UpdateNotification(&notification)
	}

	return notificationResponses, nil
}

func (s *notificationServiceImpl) GetAllUserNotification(userId string) ([]dto.NotificationResponse, error) {
	notifications, err := s.notificationRepository.GetAllUserNotification(userId)
	if err != nil {
		return nil, err
	}

	// Convert model notifications to NotificationResponse DTOs
	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			NotificationId: notification.ID.String(),
			ReservationId:  notification.ReservationID.String(),
			Content:        notification.Content,
			From:           notification.FromName,
			Time:           notification.Time,
		})
		notification.IsRead = true
		s.notificationRepository.UpdateNotification(&notification)
	}

	return notificationResponses, nil
}

func (s *notificationServiceImpl) GetAllRestaurantNotification(userId string) ([]dto.NotificationResponse, error) {
	restaurant, err := s.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil || restaurant == nil {
		return nil, err
	}
	notifications, err := s.notificationRepository.GetAllRestaurantNotification(restaurant.ID.String())
	if err != nil {
		return nil, err
	}

	// Convert model notifications to NotificationResponse DTOs
	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			NotificationId: notification.ID.String(),
			ReservationId:  notification.ReservationID.String(),
			Content:        notification.Content,
			From:           notification.FromName,
			Time:           notification.Time,
		})
		notification.IsRead = true
		s.notificationRepository.UpdateNotification(&notification)
	}

	return notificationResponses, nil
}
