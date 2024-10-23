package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkstpm/Softdev-Backend/internal/config"
	"github.com/pkstpm/Softdev-Backend/internal/users/model"
)

var (
	accessTokenSecret []byte
)

func init() {
	cfg := config.GetConfig()
	accessTokenSecret = []byte(cfg.Jwt.AccessSecret)
}

type Claims struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

func GenerateAccessToken(user *model.User) (string, error) {
	claims := Claims{
		UserID:   user.ID.String(),
		UserType: string(user.UserType),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 1 hour
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}
