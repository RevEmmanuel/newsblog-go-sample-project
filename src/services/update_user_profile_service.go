package services

import (
	"newsblog-users-ms/src/database"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
)

func UpdateUserProfile(updateFields *models.UpdateUserDTO, email string)(models.User, error) {
	var err error
	user, err := GetUserProfile(email)
	if err != nil {
		return models.User{}, err
	}
	if len(updateFields.Preferences) != 0 {
		for _, preference := range user.Preferences {
			err = database.DB.Delete(&models.Preference{}, preference.ID).Error
			if err != nil {
				break
			}
		}
		if err != nil {
			return models.User{}, err
		}
	}
	utils.UpdateStructFields(user.UserDTO, updateFields)
	// database.DB.Model().Updates()
	err = database.DB.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}