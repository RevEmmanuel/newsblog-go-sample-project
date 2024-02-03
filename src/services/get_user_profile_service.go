package services

import (
	"newsblog-users-ms/src/database"
	"newsblog-users-ms/src/models"
)

func GetUserProfile(email string) (models.User, error) {
	var user models.User
	err := database.DB.Where("email = ? AND deleted_at IS NULL", email).Preload("Preferences").First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}