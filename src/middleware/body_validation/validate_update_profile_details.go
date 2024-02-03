package middleware

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/utils"
	helpers "newsblog-users-ms/src/utils/validate_helpers"

	"github.com/gin-gonic/gin"
)

func ValidateUpdateProfileDetails(c *gin.Context) {
	var updateFields models.UpdateUserDTO
	err := c.BindJSON(&updateFields)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	err = helpers.ValidateStructBody(&updateFields)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.Set(constants.USER_UPDATE_FIELDS, &updateFields)
	c.Next()
}