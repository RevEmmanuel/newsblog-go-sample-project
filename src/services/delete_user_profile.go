package services

import (
	"newsblog-users-ms/src/database"
	"newsblog-users-ms/src/models"
)

func DeleteUserProfile(email string) error {
	user, err := GetUserProfile(email)
	if err != nil {
		return err
	}
	for _, preference := range user.Preferences {
		database.DB.Delete(&models.Preference{}, preference.ID)
	}

	err = database.DB.Delete(&models.User{}, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}
