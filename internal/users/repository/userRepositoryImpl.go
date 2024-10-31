package repository

import (
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{db: db.GetDb()}
}

func (r *userRepository) CreateUser(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Favourite").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByEmailOrUsername(identifier string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindMatchUsernameOrEmail(username string, email string) bool {
	var user model.User
	err := r.db.Where("email = ? OR username = ?", email, username).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func (r *userRepository) EditProfileUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) UpdateUserType(userId string, role string) error {
	err := r.db.Model(&model.User{}).Where("id = ?", userId).Update("user_type", role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) SearchFavouriteRestaurant(userId string, restaurantId string) bool {
	var count int64
	err := r.db.Model(&model.Favourite{}).
		Where("user_id = ? AND restaurant_id = ?", userId, restaurantId).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (r *userRepository) AddFavouriteRestaurant(favourite *model.Favourite) error {
	err := r.db.Create(favourite).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) RemoveFavouriteRestaurant(restaurantId string) error {
	err := r.db.Where("restaurant_id = ?", restaurantId).Delete(&model.Favourite{}).Error
	if err != nil {
		return err
	}
	return nil
}
