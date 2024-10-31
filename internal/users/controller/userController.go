package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	UpdateProfile(c echo.Context) error
	ViewProfile(c echo.Context) error
	ChangePassword(c echo.Context) error
	UploadUserProfilePicture(c echo.Context) error
	AddFavouriteRestaurant(c echo.Context) error
	RemoveFavouriteRestaurant(c echo.Context) error
	GetUserById(c echo.Context) error
}
