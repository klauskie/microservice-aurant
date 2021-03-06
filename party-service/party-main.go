package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/controllers"
	"klauskie.com/microservice-aurant/party-service/util"
)

func main() {
	r := gin.Default()

	r.Use(controllers.CORSMiddleware)
	r.Use(controllers.TokenAuthMiddleware)

	api := r.Group("/api")
	{
		// TODO use query params
		api.POST("/party", controllers.CreateParty)
		api.PUT("/party/:partyID", controllers.JoinParty)
		api.GET("/party/:partyID", controllers.GetParty)
		api.DELETE("/party/:partyID/kick/:clientID", controllers.KickFromParty)
		api.DELETE("/party/:partyID", controllers.RemoveParty)
		api.GET("/party-clients/:partyID", controllers.GetPartyClients)
		api.GET("/party-status/:partyID", controllers.GetPartyClients)
		api.PUT("/party-status/:partyID", controllers.UpdatePartyStatus)

		api.POST("/order/:partyID", controllers.CreateClientOrder)
		api.GET("/order/:partyID", controllers.GetClientOrder)
		api.GET("/party-order/:partyID", controllers.GetAllOrder)
		api.POST("/prepare-order/:partyID", controllers.SendPrepareCommandOrder)

		api.POST("/test/prepare-order", controllers.SendPrepareCommandOrderTest)
		api.GET("/test/item-definition", controllers.GetItemDefinition)
		api.GET("/test/party", controllers.GetAllParties)
	}

	util.ClearCacheCron()

	r.Run(":8081")
}
