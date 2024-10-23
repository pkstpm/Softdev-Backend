package controller

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	if err := c.Bind(&createDishDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&createDishDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.CreateDish(userId, &createDishDTO)
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
	var createTimeSlotDTO dto.UpdateTimeSlotDTO
	if err := c.Bind(&createTimeSlotDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := validate.Struct(&createTimeSlotDTO); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Validation failed", err.Error())
	}

	err := h.restaurantService.UpdateTimeSlot(userId, &createTimeSlotDTO)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Create TimeSlot failed", err.Error())
	}

	return utils.SendSuccess(c, "TimeSlot create successfully", nil)
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
