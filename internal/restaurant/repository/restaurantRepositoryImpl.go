package repository

import (
	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
	"gorm.io/gorm"
)

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db database.Database) RestaurantRepository {
	return &restaurantRepository{db: db.GetDb()}
}

func (r *restaurantRepository) FindRestaurantByUserID(userId string) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	err := r.db.Where("user_id = ?", userId).First(&restaurant).Error
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *restaurantRepository) FindRestaurantByID(restaurantId string) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	err := r.db.Preload("Images").Where("id = ?", restaurantId).First(&restaurant).Error
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

// FindRestaurantByName performs a case-insensitive partial match search for a restaurant by name
func (r *restaurantRepository) FindRestaurantByName(name string) ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	searchPattern := "%" + name + "%" // This will match any restaurant name containing the search term

	// Use LIKE for partial matching
	if err := r.db.Where("restaurant_name ILIKE ?", searchPattern).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// FindRestaurantByCategory performs a case-insensitive partial match search for restaurants by category
func (r *restaurantRepository) FindRestaurantByCategory(category string) ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	searchPattern := "%" + category + "%" // This will match any category containing the search term

	// Use LIKE for partial matching
	if err := r.db.Where("category ILIKE ?", searchPattern).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantRepository) CreateRestaurant(user *model.Restaurant) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) FindDishById(dishId string) (*model.Dish, error) {
	var dish model.Dish
	err := r.db.Where("id = ?", dishId).First(&dish).Error
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (r *restaurantRepository) CreateDish(dish *model.Dish) error {
	err := r.db.Create(dish).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) UpdateDish(dish *model.Dish) error {
	err := r.db.Save(dish).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) GetDishPrice(dishId uuid.UUID) (int, error) {
	var dish model.Dish
	err := r.db.Select("price").Where("id = ?", dishId).First(&dish).Error
	if err != nil {
		return 0, err
	}
	return dish.Price, nil
}

func (r *restaurantRepository) GetTimeSlotsByRestaurantId(restaurantId string) ([]model.TimeSlot, error) {
	var timeSlots []model.TimeSlot
	err := r.db.Select("id, weekday, hour_start, hour_end, restaurant_id").Where("restaurant_id = ?", restaurantId).Find(&timeSlots).Error
	if err != nil {
		return nil, err
	}
	return timeSlots, nil
}

func (r *restaurantRepository) GetTablesByRestaurantId(restaurantId string) ([]model.Table, error) {
	var tables []model.Table
	err := r.db.Select("id, table_number, capacity, restaurant_id").Where("restaurant_id = ?", restaurantId).Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *restaurantRepository) GetTableById(tableId string) (*model.Table, error) {
	var table model.Table
	err := r.db.Preload("Reservations").Where("id = ?", tableId).First(&table).Error
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *restaurantRepository) CreateTable(table *model.Table) error {
	err := r.db.Create(table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) UpdateTable(table *model.Table) error {
	err := r.db.Save(table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) DeleteTable(tableId string) error {
	err := r.db.Delete(&model.Table{}, tableId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) CreateTimeSlot(timeSlot *model.TimeSlot) error {
	err := r.db.Create(timeSlot).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) UpdateTimeSlot(timeSlot *model.TimeSlot) error {
	err := r.db.Save(timeSlot).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) CreateImages(images *model.Image) error {
	err := r.db.Create(images).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepository) DeleteImage(imageId string) error {
	err := r.db.Delete(&model.Image{}, imageId).Error
	if err != nil {
		return err
	}
	return nil
}
