package middleware

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
	helpers "newsblog-users-ms/src/utils/validate_helpers"

	"github.com/gin-gonic/gin"
)

func ValidateUserLogin(c *gin.Context) {
	var payload models.LoginPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	err = helpers.ValidateStructBody(&payload)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	
	
	c.Set(constants.LOGIN_PAYLOAD, &payload)
	c.Next()
}