package service

import (
	"github.com/pkstpm/Softdev-Backend/internal/notification/dto"
)

type NotificationService interface {
	GetUserNotReadNotification(userId string) ([]dto.NotificationResponse, error)
	GetRestaurantNotReadNotification(restaurantId string) ([]dto.NotificationResponse, error)
	GetAllUserNotification(userId string) ([]dto.NotificationResponse, error)
	GetAllRestaurantNotification(restaurantId string) ([]dto.NotificationResponse, error)
}
