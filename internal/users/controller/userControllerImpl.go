package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/users/dto"
	"github.com/pkstpm/Softdev-Backend/internal/users/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	dto.RegisterValidators(validate)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

func (h *userController) UpdateProfile(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var editProfileDTO dto.EditProfileDTO
	if err := c.Bind(&editProfileDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}
	if err := validate.Struct(&editProfileDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}
	editedUser, err := h.userService.EditProfile(userId, editProfileDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Update failed", nil)
	}
	return utils.SendSuccess(c, "User registered successfully", editedUser)
}

func (h *userController) ViewProfile(c echo.Context) error {
	userId := c.Get("user_id").(string)
	user, err := h.userService.GetUser(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to get user", nil)
	}
	return utils.SendSuccess(c, "User retrieved successfully", user)
}
