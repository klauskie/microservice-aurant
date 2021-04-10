package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/controllers"
)

func main() {
	r := gin.Default()

	r.Use(controllers.TokenAuthMiddleware)

	api := r.Group("/party-api")
	{
		// TODO use query params
		api.POST("/party/:vendorID", controllers.CreateParty)
		api.PUT("/party/:partyID", controllers.JoinParty)
		api.GET("/party/:partyID", controllers.GetParty)
		api.DELETE("/party/:partyID/kick/:client", controllers.KickFromParty)
		api.DELETE("/party/:partyID", controllers.RemoveParty)
		api.GET("/party-clients/:partyID", controllers.GetPartyClients)
		api.GET("/party-status/:partyID", controllers.GetPartyClients)
		api.PUT("/party-status/:partyID", controllers.UpdatePartyStatus)

		api.POST("/order/:partyID", controllers.CreateClientOrder)
		api.GET("/order/:partyID", controllers.GetClientOrder)
		api.GET("/party-order/:partyID", controllers.GetAllOrder)
		api.POST("/prepare-order/:partyID", controllers.SendPrepareCommandOrder)

		api.POST("/test/prepare-order", controllers.SendPrepareCommandOrderTest)
	}

	r.Run(":8081")
}
