package repository

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
)

type RestaurantRepository interface {
	FindRestaurantByUserID(userId string) (*model.Restaurant, error)
	FindRestaurantByName(name string) ([]model.Restaurant, error)
	FindRestaurantByCategory(category string) ([]model.Restaurant, error)
	CreateRestaurant(restaurant *model.Restaurant) error

	FindDishById(dishId string) (*model.Dish, error)
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish) error
	GetDishPrice(dishId uuid.UUID) (int, error)

	GetTablesByRestaurantId(restaurantId string) ([]model.Table, error)
	CreateTable(table *model.Table) error
	UpdateTable(table *model.Table) error
	DeleteTable(tableId string) error

	GetTimeSlotsByRestaurantId(restaurantId string) ([]model.TimeSlot, error)
	CreateTimeSlot(timeSlot *model.TimeSlot) error
	UpdateTimeSlot(timeSlot *model.TimeSlot) error
}
