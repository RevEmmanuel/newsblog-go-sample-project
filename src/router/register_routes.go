package router

import (
	"newsblog-users-ms/src/constants"
	"newsblog-users-ms/src/controllers"
	body "newsblog-users-ms/src/middleware/body_validation"
	header "newsblog-users-ms/src/middleware/header_validation"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.POST(constants.USER_SIGNUP_ROUTE, body.ValidateUserRegistration, controllers.UserRegistration)
	apiGroup.POST(constants.USER_LOGIN_ROUTE, body.ValidateUserLogin, controllers.UserLogin)
	
	apiGroup.GET(constants.USER_GET_PROFILE_ROUTE, header.AuthenticateUserToken, controllers.GetUserProfile)
	apiGroup.PATCH(constants.USER_UPDATE_PROFILE_ROUTE, header.AuthenticateUserToken, body.ValidateUpdateProfileDetails, controllers.UpdateUserProfile)
	apiGroup.DELETE(constants.USER_DELETE_PROFILE_ROUTE, header.AuthenticateUserToken, controllers.DeleteUserProfile)
}