package repository

import "github.com/pkstpm/Softdev-Backend/internal/users/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserByID(id string) (*model.User, error)
	FindUserByEmailOrUsername(identifier string) (*model.User, error)
	FindMatchUsernameOrEmail(username string, email string) bool
	EditProfileUser(user *model.User) error
	UpdateUserType(userId string, role string) error
}
