package controller

import "github.com/labstack/echo/v4"

type RestaurantController interface {
	GetMyRestaurant(c echo.Context) error
	GetAllRestaurants(c echo.Context) error
	FindByName(c echo.Context) error
	FindByCategory(c echo.Context) error
	CreateDish(c echo.Context) error
	GetDishesByRestaurantId(c echo.Context) error
	GetDishesById(c echo.Context) error
	UpdateDish(c echo.Context) error
	GetTimeSlot(c echo.Context) error
	GetTimeSlotById(c echo.Context) error
	UpdateTimeSlot(c echo.Context) error
	GetTable(c echo.Context) error
	CreateTable(c echo.Context) error
	UploadRestaurantPictures(c echo.Context) error
	DeleteRestauranPictures(c echo.Context) error
	GetRestaurantByID(c echo.Context) error
	GetDishByID(c echo.Context) error
	GetTableByID(c echo.Context) error
}
