package controllers

import (
	"github.com/gin-gonic/gin"
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/services"
	"newsblog-users-ms/src/utils"
)

func DeleteUserProfile(c *gin.Context) {
	value, _ := c.Get(constants.USER_TOKEN_CUSTOM_CLAIMS)
	customClaims, _ := value.(*utils.JWTClaim)
	err := services.DeleteUserProfile(customClaims.Email)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.JSON(200, utils.ResponseMessage("Delete user successful", nil))
}