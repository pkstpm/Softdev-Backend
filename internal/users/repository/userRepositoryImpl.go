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
	err := r.db.Where("id = ?", id).First(&user).Error
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
