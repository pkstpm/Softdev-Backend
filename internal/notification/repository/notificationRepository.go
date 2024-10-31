package repository

import "github.com/pkstpm/Softdev-Backend/internal/notification/model"

type NotificationRepository interface {
	GetUserNotReadNotification(userId string) ([]model.Notification, error)
	GetRestaurantNotReadNotification(restaurantId string) ([]model.Notification, error)
	GetAllUserNotification(userId string) ([]model.Notification, error)
	GetAllRestaurantNotification(restaurantId string) ([]model.Notification, error)
	CreateNotification(notification *model.Notification) error
	UpdateNotification(notification *model.Notification) error
}
