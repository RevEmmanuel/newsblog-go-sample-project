package controllers

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/services"
	"newsblog-users-ms/src/utils"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	value, _ := c.Get(constants.LOGIN_PAYLOAD)
	loginDetails, _ := value.(*models.LoginPayload)
	userDetails, accessToken, err := services.LoginUser(loginDetails)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.JSON(200, utils.ResponseMessage("Login Successful", gin.H{
		"Profile Details": userDetails,
		"Access Token": accessToken,
	}))
}