package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func CreateToken(secretKey string, id int, roleId int) (accessToken string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"user_id": id,
			"role_id": roleId,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logrus.Error("CreateToken", err)
		return
	}

	return tokenString, nil
}
