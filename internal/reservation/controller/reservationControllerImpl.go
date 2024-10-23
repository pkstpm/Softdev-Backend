package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/dto"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type reservationController struct {
	reservationService service.ReservationService
}

func NewReservationController(reservationService service.ReservationService) ReservationController {
	return &reservationController{reservationService: reservationService}
}

func (h *reservationController) CreateReservation(c echo.Context) error {
	userIdStr := c.Get("user_id").(string)

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid user ID format", err.Error())
	}

	var createReservationDTO dto.CreateReservationDTO
	if err := c.Bind(&createReservationDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&createReservationDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err = h.reservationService.CreateReservation(userId, createReservationDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create reservation failed", err.Error())
	}

	return utils.SendSuccess(c, "Reservation created successfully", nil)
}
