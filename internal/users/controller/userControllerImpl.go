package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (h *userController) ChangePassword(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var changePasswordDTO dto.ChangePasswordDTO

	if err := c.Bind(&changePasswordDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}
	if err := validate.Struct(&changePasswordDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.userService.ChangePassword(userId, changePasswordDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to change password", err.Error())
	}

	return utils.SendSuccess(c, "Password changed successfully", nil)
}

func (h *userController) UploadUserProfilePicture(c echo.Context) error {
	userId := c.Get("user_id").(string)

	// Retrieve the file from the form
	file, err := c.FormFile("image")
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Failed to retrieve profile picture", err)
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to open profile picture", err)
	}
	defer src.Close()

	// Check the MIME type of the file
	mimeType := file.Header.Get("Content-Type")
	if !utils.IsAllowedMimeType(mimeType) {
		return utils.SendError(c, http.StatusBadRequest, "Invalid file type. Only PNG, JPEG, or JPG are allowed.", nil)
	}

	// Create a unique file name using UUID
	uniqueID := uuid.New().String()
	fileExtension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", uniqueID, fileExtension)

	// Define the destination path
	dstPath := filepath.Join("uploads", newFileName)

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create upload directory", err)
	}

	// Create the destination file
	dst, err := os.Create(dstPath)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create destination file", err)
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to upload profile picture", err)
	}

	// Call your service to handle the upload
	user, err := h.userService.UploadUserProfilePicture(userId, dstPath)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to upload profile picture", err)
	}

	// Return success response
	return utils.SendSuccess(c, "Profile picture uploaded successfully", user.ImgPath)
}

func (h *userController) AddFavouriteRestaurant(c echo.Context) error {
	userId := c.Get("user_id").(string)
	restaurantId := c.Param("restaurant_id")

	err := h.userService.AddFavouriteRestaurant(userId, restaurantId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to add favourite restaurant", err.Error())
	}

	return utils.SendSuccess(c, "Favourite restaurant added successfully", nil)
}

func (h *userController) RemoveFavouriteRestaurant(c echo.Context) error {
	userId := c.Get("user_id").(string)
	restaurantId := c.Param("restaurant_id")

	err := h.userService.RemoveFavouriteRestaurant(userId, restaurantId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to remove favourite restaurant", err.Error())
	}

	return utils.SendSuccess(c, "Favourite restaurant removed successfully", nil)
}
