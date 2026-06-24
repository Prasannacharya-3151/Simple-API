package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	claims := Claims{
		UserID: userID,
		Email:email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := &Claims{} //which wil creates a enpty struct Claims{ USerID:0, Email:"",}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid{
		return nil, err
	}

	return claims, nil
}