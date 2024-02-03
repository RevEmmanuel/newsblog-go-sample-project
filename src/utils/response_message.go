package utils

import "newsblog-users-ms/src/models"

func ResponseMessage(message string, data any) models.ResponseMessage{
	return models.ResponseMessage{
		Message: message,
		Data: data,
	}
}