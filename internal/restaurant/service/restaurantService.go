package service

import (
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/dto"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
)

type RestaurantService interface {
	GetAllRestaurants() ([]model.Restaurant, error)
	GetAllDishesByRestaurantId(restaurantId string) ([]model.Dish, error)
	FindRestaurantByName(name string) ([]model.Restaurant, error)
	FindRestaurantByCategory(category string) ([]model.Restaurant, error)
	CreateDish(userId string, dto *dto.CreateDishDTO) error
	UpdateDish(userId string, dto *dto.UpdateDishDTO) error
	GetTimeSlot(userId string) ([]model.TimeSlot, error)
	GetTimeSlotByRestaurantId(restaurantId string) ([]model.TimeSlot, error)
	CreateTimeSlot(userId string) error
	UpdateTimeSlot(userId string, timeSlot *dto.UpdateTimeDTO) error
	GetTablesByRestaurantId(restaurantId string) ([]model.Table, error)
	GetRestaurantByUserId(userId string) (*model.Restaurant, error)
	CreateTable(userId string, dto *dto.CreateTableDTO) error
	UploadRestaurantPictures(userId string, uploadedFiles []string) error
	DeletetRestaurantPicture(userId string, pictureId string) error
	GetRestaurantByID(restaurantId string) (*model.Restaurant, error)
}
