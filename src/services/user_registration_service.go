package services

import (
	"newsblog-users-ms/src/database"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
)

func RegisterUser(newUser *models.User) (string, error) {
	err := utils.HashPassword(newUser)
	if err != nil {
		return "", err
	}
	err = database.DB.Create(newUser).Error
	if err != nil {
		return "", err
	}
	customClaims := utils.JWTClaim{
		Role: "user",
		Email: newUser.Email,
	}
	accessToken, err := customClaims.GenenerateAccessToken()
	if err != nil {
		return "", err
	}
	return accessToken, nil 
	
}