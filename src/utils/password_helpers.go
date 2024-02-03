package utils

import (
	"newsblog-users-ms/src/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(NewUser *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(NewUser.Password), 10)
	if err != nil {
		return err
	}
	NewUser.Password = string(bytes)
	return nil
}
func CheckPassword(actualPassword string, providedPassword string) error  {
	err := bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
