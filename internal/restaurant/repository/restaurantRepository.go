package repository

import (
	"github.com/google/uuid"
	reservationModel "github.com/pkstpm/Softdev-Backend/internal/reservation/model"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
)

type RestaurantRepository interface {
	GetAllRestaurants() ([]model.Restaurant, error)
	FindRestaurantByUserID(userId string) (*model.Restaurant, error)
	FindRestaurantByID(restaurantId string) (*model.Restaurant, error)
	FindRestaurantByName(name string) ([]model.Restaurant, error)
	FindRestaurantByCategory(category string) ([]model.Restaurant, error)
	CreateRestaurant(restaurant *model.Restaurant) error

	AddReservationToTable(tableId string, reservation *reservationModel.Reservation) error

	UpdateRestaurant(restaurant *model.Restaurant) error

	FindDishById(dishId string) (*model.Dish, error)
	FindDishByName(name string, restauranId string) (*model.Dish, error)
	GetAllDishesByRestaurantId(restaurantId string) ([]model.Dish, error)
	CreateDish(dish *model.Dish) error
	GetDishesByRestaurantId(restaurantId string) ([]model.Dish, error)
	UpdateDish(dish *model.Dish) error
	GetDishPrice(dishId uuid.UUID) (int, error)

	GetTablesByRestaurantId(restaurantId string) ([]model.Table, error)
	GetTableById(tableId string) (*model.Table, error)
	CreateTable(table *model.Table) error
	UpdateTable(table *model.Table) error
	DeleteTable(tableId string) error

	GetTimeSlotsByRestaurantId(restaurantId string) ([]model.TimeSlot, error)
	CreateTimeSlot(timeSlot *model.TimeSlot) error
	UpdateTimeSlot(timeSlot *model.TimeSlot) error

	CreateImages(images *model.Image) error
	DeleteImage(imageId uuid.UUID) error

	GetDishByID(dishId string) (*model.Dish, error)
	GetTableByID(tableId string) (*model.Table, error)
}
