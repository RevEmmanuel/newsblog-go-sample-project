package controllers

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/models"
	"newsblog-users-ms/src/services"
	"newsblog-users-ms/src/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfile(c *gin.Context) {
	value, _ := c.Get(constants.USER_TOKEN_CUSTOM_CLAIMS)
	customClaims, _ := value.(*utils.JWTClaim)
	value, _ = c.Get(constants.USER_UPDATE_FIELDS)
	updateFields, _ := value.(*models.UpdateUserDTO)
	updatedUserDetails, err := services.UpdateUserProfile(updateFields, customClaims.Email)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.JSON(200, utils.ResponseMessage("User details updated successfully", updatedUserDetails))
}