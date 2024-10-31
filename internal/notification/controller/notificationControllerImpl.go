package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/notification/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

type notificationController struct {
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) NotificationController {
	return &notificationController{notificationService: notificationService}
}

func (h *notificationController) GetUserNotReadNotification(c echo.Context) error {
	userId := c.Get("user_id").(string)
	notifications, err := h.notificationService.GetUserNotReadNotification(userId)
	if err != nil {
		return utils.SendError(c, 500, "Get user not read notification failed", err.Error())
	}
	return utils.SendSuccess(c, "Get user not read notification success", notifications)
}

func (h *notificationController) GetRestaurantNotReadNotification(c echo.Context) error {
	restaurantId := c.Get("user_id").(string)
	notifications, err := h.notificationService.GetRestaurantNotReadNotification(restaurantId)
	if err != nil {
		return utils.SendError(c, 500, "Get restaurant not read notification failed", err.Error())
	}
	return utils.SendSuccess(c, "Get restaurant not read notification success", notifications)
}

func (h *notificationController) GetAllUserNotification(c echo.Context) error {
	userId := c.Get("user_id").(string)
	notifications, err := h.notificationService.GetAllUserNotification(userId)
	if err != nil {
		return utils.SendError(c, 500, "Get all user notification failed", err.Error())
	}
	return utils.SendSuccess(c, "Get all user notification success", notifications)
}

func (h *notificationController) GetAllRestaurantNotification(c echo.Context) error {
	restaurantId := c.Get("user_id").(string)
	notifications, err := h.notificationService.GetAllRestaurantNotification(restaurantId)
	if err != nil {
		return utils.SendError(c, 500, "Get all restaurant notification failed", err.Error())
	}
	return utils.SendSuccess(c, "Get all restaurant notification success", notifications)
}
