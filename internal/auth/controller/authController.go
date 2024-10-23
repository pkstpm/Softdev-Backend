package controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	RegisterRestaurant(c echo.Context) error
	Me(c echo.Context) error
}
