package middleware

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
	helpers "newsblog-users-ms/src/utils/validate_helpers"

	"github.com/gin-gonic/gin"
)

func ValidateUserRegistration(c *gin.Context) {
	var reqBody models.UserDTO
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	err = helpers.ValidateStructBody(&reqBody)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	var NewUser = models.User{
		UserDTO: &reqBody,
	}
	c.Set(constants.NEW_USER_OBJECT, &NewUser)
	c.Next()

}
