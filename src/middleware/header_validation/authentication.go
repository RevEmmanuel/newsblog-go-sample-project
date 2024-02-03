package middleware

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateUserToken(c *gin.Context) {
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		c.AbortWithStatusJSON(401, utils.ResponseMessage("No access token provided", nil))
		return
	}
	seperatedToken := strings.Split(clientToken, "Bearer ")
	if len(seperatedToken) != 2 {
		c.AbortWithStatusJSON(401, utils.ResponseMessage("Invalid access token", nil))
		return
	}
	signedToken := seperatedToken[1]
	customClaims, err := utils.ValidateToken(signedToken)
	if err != nil {
		c.AbortWithStatusJSON(401, utils.ResponseMessage(err.Error(), nil))
		return
	}
	c.Set(constants.USER_TOKEN_CUSTOM_CLAIMS, customClaims)
	c.Next()
}