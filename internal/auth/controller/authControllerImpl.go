package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/auth/dto"
	"github.com/pkstpm/Softdev-Backend/internal/auth/service"
	restaurantService "github.com/pkstpm/Softdev-Backend/internal/restaurant/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	dto.RegisterValidators(validate)
}

type authController struct {
	authService       service.AuthService
	restaurantService restaurantService.RestaurantService
}

func NewAuthController(authService service.AuthService, restaurantService restaurantService.RestaurantService) AuthController {
	return &authController{authService: authService, restaurantService: restaurantService}
}

func (h *authController) Register(c echo.Context) error {
	var registerDTO dto.RegisterDTO
	if err := c.Bind(&registerDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}
	if err := validate.Struct(&registerDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.authService.Register(&registerDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Registration failed", err.Error())
	}

	return utils.SendSuccess(c, "User registered successfully", nil)
}

func (h *authController) Login(c echo.Context) error {
	var loginDTO dto.LoginDTO
	if err := c.Bind(&loginDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&loginDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	accessToken, err := h.authService.Login(&loginDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Login failed", nil)
	}

	return utils.SendSuccess(c, "User login successfully", accessToken)
}

func (h *authController) RegisterRestaurant(c echo.Context) error {
	userIdStr := c.Get("user_id").(string)
	var registerRestaurantDTO dto.RegisterRestaurantDTO
	if err := c.Bind(&registerRestaurantDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid user ID format", err.Error())
	}

	if err := validate.Struct(&registerRestaurantDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err = h.authService.RegisterRestaurant(userId, &registerRestaurantDTO)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return utils.SendError(c, http.StatusConflict, "Restaurant already exists for this user", nil)
		}
		return utils.SendError(c, http.StatusInternalServerError, "Registration failed", err.Error())
	}

	err = h.restaurantService.CreateTimeSlot(userId.String())
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create Time Slot Error", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant registered successfully", nil)
}

func (h *authController) Me(c echo.Context) error {
	userIdStr := c.Get("user_id").(string)

	user, accessToken, err := h.authService.Me(userIdStr)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Login failed", nil)
	}

	var favourites []dto.FavoriteRestaurantDTO
	for _, fav := range user.Favourite {
		favourites = append(favourites, dto.FavoriteRestaurantDTO{
			RestaurantID: fav.RestaurantID.String(),
		})
		log.Printf(fav.RestaurantID.String())
	}

	userResponse := dto.UserResponse{
		ID:          user.ID.String(),
		DisplayName: user.DisplayName,
		Username:    user.Username,
		Email:       user.Email,
		Role:        string(user.UserType),
		ImgPath:     user.ImgPath,
		PhoneNumebr: user.PhoneNumber,
		Favourite:   favourites,
		AccessToken: accessToken,
	}

	return utils.SendSuccess(c, "OK", userResponse)
}
