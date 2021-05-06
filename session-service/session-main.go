package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/session-service/controllers"
	"klauskie.com/microservice-aurant/session-service/util"
)

func main() {
	r := gin.Default()

	r.Use(controllers.CORSMiddleware)

	api := r.Group("/api")
	{
		// TODO use query params
		api.POST("/register", controllers.RegisterClient)
		api.POST("/login", controllers.Login)
		api.POST("/guest-login", controllers.LoginAsGuest)
		api.DELETE("/logout", controllers.Logout)
		api.GET("/profile", controllers.Profile)
		api.GET("/token-validation/:token", controllers.TokenValidation)
	}

	util.ClearCacheCron()

	r.Run(":8083")
}
