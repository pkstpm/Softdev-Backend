package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/dto"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/service"
	"github.com/pkstpm/Softdev-Backend/internal/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type restaurantController struct {
	restaurantService service.RestaurantService
}

func NewRestaurantController(restaurantService service.RestaurantService) RestaurantController {
	return &restaurantController{restaurantService: restaurantService}
}

func (h *restaurantController) FindByName(c echo.Context) error {
	name := c.QueryParam("name")
	restaurants, err := h.restaurantService.FindRestaurantByName(name)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Find restaurant by name failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant retrieved successfully", restaurants)
}

func (h *restaurantController) FindByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	restaurants, err := h.restaurantService.FindRestaurantByCategory(category)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Find restaurant by category failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant retrieved successfully", restaurants)
}
func (h *restaurantController) CreateDish(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var createDishDTO dto.CreateDishDTO

	// Manually get form values for the fields
	createDishDTO.Name = c.FormValue("name")
	createDishDTO.Description = c.FormValue("description")
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid price format", nil)
	}
	createDishDTO.Price = price

	// Validate the DTO after assigning the values
	if err := validate.Struct(&createDishDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	// Retrieve the file
	form, err := c.MultipartForm()
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Failed to retrieve files", err)
	}

	files := form.File["image"]
	if len(files) == 0 {
		return utils.SendError(c, http.StatusBadRequest, "No image uploaded", nil)
	}

	file := files[0]
	src, err := file.Open()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to open file", err)
	}
	defer src.Close()

	mimeType := file.Header.Get("Content-Type")
	if !utils.IsAllowedMimeType(mimeType) {
		return utils.SendError(c, http.StatusBadRequest, "Invalid file type. Only PNG, JPEG, or JPG are allowed.", nil)
	}

	uniqueID := uuid.New().String()
	fileExtension := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", uniqueID, fileExtension)
	dstPath := filepath.Join("uploads", newFileName)

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create upload directory", err)
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create destination file", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to upload file", err)
	}

	err = h.restaurantService.CreateDish(userId, &createDishDTO, dstPath)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create dish failed", err.Error())
	}

	return utils.SendSuccess(c, "Dish created successfully", nil)
}

func (h *restaurantController) UpdateDish(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var updateDishDTO dto.UpdateDishDTO
	if err := c.Bind(&updateDishDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&updateDishDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.UpdateDish(userId, &updateDishDTO)

	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Update dish failed", err.Error())
	}

	return utils.SendSuccess(c, "Dish updated successfully", nil)
}

func (h *restaurantController) GetTimeSlot(c echo.Context) error {
	userId := c.Get("user_id").(string)
	timeSlots, err := h.restaurantService.GetTimeSlot(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Get TimeSlot failed", err.Error())
	}

	return utils.SendSuccess(c, "TimeSlot retrieved successfully", timeSlots)
}

func (h *restaurantController) UpdateTimeSlot(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var updateTimeDTO dto.UpdateTimeDTO
	if err := c.Bind(&updateTimeDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&updateTimeDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.UpdateTimeSlot(userId, &updateTimeDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create TimeSlot failed", err.Error())
	}

	return utils.SendSuccess(c, "TimeSlot update successfully", nil)
}

func (h *restaurantController) GetTable(c echo.Context) error {
	id := c.Param("restaurant_id")
	log.Print(id)
	tables, err := h.restaurantService.GetTablesByRestaurantId(id)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Find table by restaurantid failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant retrieved successfully", tables)
}

func (h *restaurantController) CreateTable(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var createTableDTO dto.CreateTableDTO
	if err := c.Bind(&createTableDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&createTableDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.CreateTable(userId, &createTableDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create Table failed", err.Error())
	}

	return utils.SendSuccess(c, "Table created successfully", nil)
}

func (h *restaurantController) UploadRestaurantPictures(c echo.Context) error {
	userId := c.Get("user_id").(string)

	// Retrieve the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Failed to retrieve files", err)
	}

	// Retrieve multiple files from the form
	files := form.File["images"]
	if len(files) == 0 {
		return utils.SendError(c, http.StatusBadRequest, "No files uploaded", nil)
	}

	var uploadedFiles []string // Store paths of successfully uploaded files

	for _, file := range files {
		// Open the file
		src, err := file.Open()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to open file", err)
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
			return utils.SendError(c, http.StatusInternalServerError, "Failed to upload file", err)
		}

		// Optionally, you can store the uploaded file path in a slice or database
		uploadedFiles = append(uploadedFiles, dstPath)
	}

	// Call your service to handle the uploaded files if needed
	err = h.restaurantService.UploadRestaurantPictures(userId, uploadedFiles)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to upload profile pictures", err)
	}

	// Return success response with all uploaded file paths
	return utils.SendSuccess(c, "Profile pictures uploaded successfully", uploadedFiles)
}

func (h *restaurantController) GetRestaurantByID(c echo.Context) error {
	id := c.Param("restaurant_id")
	restaurant, err := h.restaurantService.GetRestaurantByID(id)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Find restaurant by id failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant retrieved successfully", restaurant)
}

func (h *restaurantController) DeleteRestauranPictures(c echo.Context) error {
	userId := c.Get("user_id").(string)
	imageId := c.Param("image_id")
	err := h.restaurantService.DeletetRestaurantPicture(userId, imageId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Find restaurant by id failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant retrieved successfully", nil)
}

func (h *restaurantController) GetTimeSlotById(c echo.Context) error {
	restaurantid := c.Param("restaurant_id")
	timeSlots, err := h.restaurantService.GetTimeSlotByRestaurantId(restaurantid)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Get TimeSlot failed", err.Error())
	}

	return utils.SendSuccess(c, "TimeSlot retrieved successfully", timeSlots)
}

func (h *restaurantController) GetAllRestaurants(c echo.Context) error {
	restaurants, err := h.restaurantService.GetAllRestaurants()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	return utils.SendSuccess(c, "Get All Restarants Success", restaurants)
}

func (h *restaurantController) GetDishesByRestaurantId(c echo.Context) error {
	restaurantId := c.Param("restaurant_id")
	dishes, err := h.restaurantService.GetAllDishesByRestaurantId(restaurantId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	return utils.SendSuccess(c, "Get Dishes By Restaurant Id Success", dishes)
}

func (h *restaurantController) GetDishesById(c echo.Context) error {
	userId := c.Get("user_id").(string)
	restaurant, err := h.restaurantService.GetRestaurantByUserId(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	dishes, err := h.restaurantService.GetAllDishesByRestaurantId(restaurant.ID.String())
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to get dishes", err.Error())
	}
	if len(dishes) == 0 {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", "No dishes found")
	}

	return utils.SendSuccess(c, "Get Dishes By User Id Success", dishes)
}

func (h *restaurantController) GetMyRestaurant(c echo.Context) error {
	userId := c.Get("user_id").(string)
	restaurant, err := h.restaurantService.GetRestaurantByUserId(userId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	return utils.SendSuccess(c, "Get Restaurant By User Id Success", restaurant)
}

func (h *restaurantController) GetDishByID(c echo.Context) error {
	dishId := c.Param("dish_id")
	dish, err := h.restaurantService.GetDishByID(dishId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	return utils.SendSuccess(c, "Get Dish By ID Success", dish)
}

func (h *restaurantController) GetTableByID(c echo.Context) error {
	tableId := c.Param("table_id")
	table, err := h.restaurantService.GetTableByID(tableId)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	return utils.SendSuccess(c, "Get Table By ID Success", table)
}

func (h *restaurantController) UpdateRestaurant(c echo.Context) error {
	userId := c.Get("user_id").(string)
	var updateRestaurantDTO dto.UpdateRestaurantDTO
	if err := c.Bind(&updateRestaurantDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&updateRestaurantDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.UpdateRestaurant(userId, &updateRestaurantDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Update restaurant failed", err.Error())
	}

	return utils.SendSuccess(c, "Restaurant updated successfully", nil)
}
