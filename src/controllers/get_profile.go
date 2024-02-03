package controllers

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/services"
	"newsblog-users-ms/src/utils"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	value, _ := c.Get(constants.USER_TOKEN_CUSTOM_CLAIMS)
	customClaims, _ := value.(*utils.JWTClaim)
	userDetails, err := services.GetUserProfile((*customClaims).Email)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.JSON(200, utils.ResponseMessage("Get Profile Successful", userDetails))
}