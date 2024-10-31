package repository

import (
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/notification/model"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db database.Database) NotificationRepository {
	return &notificationRepository{db: db.GetDb()}
}

func (n *notificationRepository) GetUserNotReadNotification(userId string) ([]model.Notification, error) {
	var notifications []model.Notification
	err := n.db.Where("user_id = ? AND is_read = false AND receiver = ?", userId, "user").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *notificationRepository) GetRestaurantNotReadNotification(restaurantId string) ([]model.Notification, error) {
	var notifications []model.Notification
	err := n.db.Where("restaurant_id = ? AND is_read = false AND receiver = ?", restaurantId, "restaurant").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *notificationRepository) CreateNotification(notification *model.Notification) error {
	err := n.db.Create(notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (n *notificationRepository) GetAllUserNotification(userId string) ([]model.Notification, error) {
	var notifications []model.Notification
	err := n.db.Where("user_id = ? AND receiver = ?", userId, "user").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *notificationRepository) GetAllRestaurantNotification(restaurantId string) ([]model.Notification, error) {
	var notifications []model.Notification
	err := n.db.Where("restaurant_id = ? AND receiver = ?", restaurantId, "restaurant").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *notificationRepository) UpdateNotification(notification *model.Notification) error {
	err := n.db.Save(notification).Error
	if err != nil {
		return err
	}
	return nil
}
