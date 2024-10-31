package controller

import "github.com/labstack/echo/v4"

type NotificationController interface {
	GetUserNotReadNotification(c echo.Context) error
	GetRestaurantNotReadNotification(c echo.Context) error
	GetAllUserNotification(c echo.Context) error
	GetAllRestaurantNotification(c echo.Context) error
}
