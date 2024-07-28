package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
)

func VerifyToken(tokenString string) (bool, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			log.Println("Invalid signing method")
			return nil, errors.New("invalid signing method")
		}

		return jwtSecret, nil
	})
	if err != nil {
		log.Println("Failed to parse token: ", err)
		return false, "", err
	}

	if !token.Valid {
		log.Println("Invalid token")
		return false, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Failed to get claims")
		return false, "", err
	}

	return true, claims["id"].(string), nil
}
