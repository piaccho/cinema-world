package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"piaccho/cinema-api/configs"
	"time"
)

func CreateToken(userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": userEmail,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(configs.EnvJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
