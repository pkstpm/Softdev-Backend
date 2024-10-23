package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response is a standard response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendSuccess sends a success response
func SendSuccess(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// SendError sends an error response
func SendError(c echo.Context, code int, message string, data interface{}) error {
	c.Logger().Errorf("Error %d: %s", code, message)
	return c.JSON(code, Response{
		Status:  "error",
		Message: message,
		Data:    data,
	})
}
