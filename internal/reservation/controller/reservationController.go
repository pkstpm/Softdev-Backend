package controller

import "github.com/labstack/echo/v4"

type ReservationController interface {
	GetReservationByUserId(c echo.Context) error
	GetReservationById(c echo.Context) error
	CreateReservation(c echo.Context) error
	AddDishItem(c echo.Context) error
}
