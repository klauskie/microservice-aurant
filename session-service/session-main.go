package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/session-service/controllers"
)

func main() {
	r := gin.Default()

	api := r.Group("/session-api")
	{
		// TODO use query params
		api.POST("/register", controllers.RegisterClient)
		api.POST("/login", controllers.Login)
		api.POST("/guest-login", controllers.LoginAsGuest)
		api.DELETE("/logout", controllers.Logout)
		api.GET("/profile", controllers.Profile)
	}

	r.Run(":8083")
}
