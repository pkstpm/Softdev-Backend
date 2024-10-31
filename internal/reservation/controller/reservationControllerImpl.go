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

	reservationId, err := h.reservationService.CreateReservation(userId, createReservationDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create reservation failed", err.Error())
	}

	return utils.SendSuccess(c, "Reservation created successfully", reservationId)
}

func (h *reservationController) GetReservationById(c echo.Context) error {
	reservationId := c.Param("reservation_id")

	reservation, err := h.reservationService.GetReservationById(reservationId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Get reservation by ID failed", err.Error())
	}

	return utils.SendSuccess(c, "Reservation retrieved successfully", reservation)
}

func (h *reservationController) GetReservationByUserId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	reservations, err := h.reservationService.GetReservationsByUserId(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Get reservations by user ID failed", err.Error())
	}

	return utils.SendSuccess(c, "Reservations retrieved successfully", reservations)
}

func (h *reservationController) AddDishItem(c echo.Context) error {
	userId := c.Get("user_id").(string)
	reservationId := c.Param("reservation_id")

	var addDishItemDTO dto.AddDishItemDTO

	if err := c.Bind(&addDishItemDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&addDishItemDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.reservationService.AddDishItem(userId, reservationId, addDishItemDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Add dish item failed", err.Error())
	}

	return utils.SendSuccess(c, "Dish item added successfully", nil)

}

func (h *reservationController) GetReservationByRestaurantId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	reservations, err := h.reservationService.GetReservationsByRestaurantId(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Get reservations by restaurant ID failed", err.Error())
	}

	return utils.SendSuccess(c, "Reservations retrieved successfully", reservations)
}
