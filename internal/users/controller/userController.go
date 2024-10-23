package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	UpdateProfile(c echo.Context) error
	ViewProfile(c echo.Context) error
}
