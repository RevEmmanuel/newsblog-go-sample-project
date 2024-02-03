package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func ValidateToken(signedToken string) (*JWTClaim, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(*jwt.Token) (interface{}, error){
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	customClaims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse token claims")
	}
	if customClaims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired by %d", (time.Now().Unix() - customClaims.ExpiresAt))
	}
	return customClaims, nil

}