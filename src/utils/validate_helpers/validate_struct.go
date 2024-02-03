package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStructBody(reqBody interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("validPreferences", ValidateUserPreferences)
	err := validate.Struct(reqBody)
	if err != nil {
		var msg string
		if v, ok := err.(validator.ValidationErrors); ok {
			for _, err := range v {
				msg = "Field " + err.Field() + " failed on the " + err.Tag() + " tag." 
				break
			}
			return errors.New(msg)
		} else {
			return err
		}
	}
	return nil

}