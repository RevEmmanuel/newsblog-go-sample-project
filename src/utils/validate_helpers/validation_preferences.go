package utils

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"

	"github.com/go-playground/validator/v10"
)

func ValidateUserPreferences(f1 validator.FieldLevel) bool  {
	userPreferences, ok := f1.Field().Interface().([]models.Preference)
	if !ok {
		return false
	}
	for _, preference := range userPreferences {
		if _, exists := constants.AllowedPreferences[preference.Name]; !exists {
			return false
		}
	}
	return true

}
