package controllers

import (
	// "newsblog-users-ms/src/models"

	"log"
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/services"
	"newsblog-users-ms/src/utils"

	"github.com/gin-gonic/gin"
)

func UserRegistration(c *gin.Context) {
	value, _ := c.Get(constants.NEW_USER_OBJECT)
	NewUser, _ := value.(*models.User)
	accessToken, err := services.RegisterUser(NewUser)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.JSON(200, utils.ResponseMessage("Successfully created user", gin.H{
		"AccessToken": accessToken,
	}))
}