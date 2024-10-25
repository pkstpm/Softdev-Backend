package controller

import "github.com/labstack/echo/v4"

type RestaurantController interface {
	FindByName(c echo.Context) error
	FindByCategory(c echo.Context) error
	CreateDish(c echo.Context) error
	UpdateDish(c echo.Context) error
	GetTimeSlot(c echo.Context) error
	UpdateTimeSlot(c echo.Context) error
	GetTable(c echo.Context) error
	CreateTable(c echo.Context) error
	UploadRestaurantPictures(c echo.Context) error
	DeleteRestauranPictures(c echo.Context) error
	GetRestaurantByID(c echo.Context) error
}
