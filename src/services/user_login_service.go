package services

import (
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
)

func LoginUser(loginDetails *models.LoginPayload) (models.User, string, error) {
	user, err := GetUserProfile(loginDetails.Email)
	if err != nil {
		return models.User{}, "", err
	}
	err = utils.CheckPassword(user.Password, loginDetails.Password)
	if err != nil {
		return models.User{}, "", err
	}
	customClaims := utils.JWTClaim{
		Email: loginDetails.Email,
		Role: "user",
	}
	accessToken, err := customClaims.GenenerateAccessToken()
	if err != nil {
		return models.User{}, "", err
	}
	return user, accessToken, err
}