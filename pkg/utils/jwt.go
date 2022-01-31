package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func GenerateJWT(secret string, userId int64) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	return claims.SignedString([]byte(secret))

}

func ValidateJWT(token, secret string) (string, error) {

	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		return claims["iss"].(string), nil
	} else {
		return "", err
	}

}