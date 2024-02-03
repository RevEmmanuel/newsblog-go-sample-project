package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTClaim struct {
	Role  string
	Email string
	jwt.StandardClaims
}

func (customClaims *JWTClaim) GenenerateAccessToken()(string, error){
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	expirationMins, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_MINS"))
	secretKey := os.Getenv("SECRET_KEY")
	customClaims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(expirationMins)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}