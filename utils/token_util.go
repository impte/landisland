package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(key string, value string, secret string) (string, error) {
	if StringIsEmpty(key) || StringIsEmpty(value) || StringIsEmpty(secret) {
		return "", errors.New("参数缺失")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key: value,
		"expire": time.Now().Unix(),
	})
	return token.SignedString([]byte(secret))
}
