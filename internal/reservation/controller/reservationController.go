package controller

import "github.com/labstack/echo/v4"

type ReservationController interface {
	CreateReservation(c echo.Context) error
}
