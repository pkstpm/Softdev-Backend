package controller

import "github.com/labstack/echo/v4"

type ReviewController interface {
	CreateReview(c echo.Context) error
	// GetReviewByReservationId(c echo.Context) error
}
