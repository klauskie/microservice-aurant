package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/controllers"
)

func main() {
	r := gin.Default()

	// TODO handle client
	// TODO handle validations

	api := r.Group("/party-api")
	{
		api.POST("/party/:vendorID", controllers.CreateParty)
		api.PUT("/party/:partyID", controllers.JoinParty)
		api.GET("/party/:partyID", controllers.GetParty)
		api.DELETE("/party/:partyID", controllers.RemoveParty)

		api.POST("/order/:partyID", controllers.CreateClientOrder)
		api.GET("/order/:partyID", controllers.GetClientOrder)
		api.GET("/party-order/:partyID", controllers.GetAllOrder)
	}

	r.Run(":8081")
}
