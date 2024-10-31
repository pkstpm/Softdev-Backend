package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/review/dto"
	"github.com/pkstpm/Softdev-Backend/internal/review/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type reviewController struct {
	reviewService service.ReviewService
}

func NewReviewController(reviewService service.ReviewService) ReviewController {
	return &reviewController{reviewService: reviewService}
}

func (h *reviewController) CreateReview(c echo.Context) error {
	userId := c.Get("user_id").(string)
	reservationID := c.Param("reservation-id")
	var reviewDTO dto.ReviewDTO
	if err := c.Bind(&reviewDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&reviewDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	err := h.reviewService.CreateReview(userId, reservationID, reviewDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create review failed", err.Error())
	}

	return utils.SendSuccess(c, "Review created successfully", nil)
}
