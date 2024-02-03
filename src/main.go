package main

import (
	"newsblog-users-ms/src/database"
	"newsblog-users-ms/src/router"
	"newsblog-users-ms/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	database.InitDatabase()
	
	
	apiGroup := app.Group("/api")
	router.RegisterRoutes(apiGroup)
	
	app.Use(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, utils.ResponseMessage("Error with server", nil))
	})

	app.Run(":8080")

}
