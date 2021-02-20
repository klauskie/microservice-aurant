package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/microservice-aurant/party-service/models"
	"klauskie.com/microservice-aurant/party-service/repository"
)

// GET /party/:vendorID
func CreateParty(c *gin.Context) {
	vendorUUID := c.Param("vendorID")

	party := models.NewParty(vendorUUID, models.Client{})
	repository.GetPartyRepository().Add(party)

	c.JSON(201, gin.H{
		"message": "Party created successfully",
		"party-tag": party.TAG,
		"party": party,
	})
}

// GET /party/:partyID
func GetParty(c *gin.Context) {
	partyId := c.Param("partyID")
	party := repository.GetPartyRepository().Get(partyId)

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"party": party,
	})
}

// PUT /party/:partyID
func JoinParty(c *gin.Context) {
	partyId := c.Param("partyID")
	client := models.Client{}

	party := repository.GetPartyRepository().Get(partyId)
	party.AddClient(client)

	c.JSON(200, gin.H{
		"message": "Joined to party successfully",
		"party-tag": party.TAG,
		"party": party,
	})
}

// DELETE /party/:partyID/kick/:client
func KickFromParty(c *gin.Context) {
	partyId := c.Param("partyID")
	client := models.Client{}

	party := repository.GetPartyRepository().Get(partyId)
	party.RemoveClient(client)

	c.JSON(202, gin.H{
		"message": "DELETE client",
	})
}

// DELETE /party/:partyID
func RemoveParty(c *gin.Context) {
	partyId := c.Param("partyID")

	repository.GetPartyRepository().Get(partyId)
	// TODO upload party order

	err := repository.GetPartyRepository().Remove(partyId)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(202, gin.H{
		"message": "DELETE party",
	})
}

// GET /party-clients/:partyID
func GetPartyClients(c *gin.Context) {
	partyId := c.Param("partyID")
	party := repository.GetPartyRepository().Get(partyId)

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"clients": party.GetClients,
	})
}

// GET /party-status/:partyID
func GetPartyStatus(c *gin.Context) {
	partyId := c.Param("partyID")
	party := repository.GetPartyRepository().Get(partyId)

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"status": party.IsOk,
	})
}

// PUT /party-status/:partyID
func UpdatePartyStatus(c *gin.Context) {
	partyId := c.Param("partyID")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	status := struct {
		ready bool
	}{}
	if err := json.Unmarshal(body, &status); err != nil {
		fmt.Println(err.Error())
	}

	party := repository.GetPartyRepository().Get(partyId)

	// TODO Validate sender is host?
	client := models.Client{}
	if party.GetHost().Id != client.Id {
		c.JSON(400, gin.H{
			"message": "Action must be performed by party host",
			"party-tag": party.TAG,
		})
		return
	}

	party.IsOk = status.ready

	c.JSON(202, gin.H{
		"message": "Party status updated",
		"party-tag": party.TAG,
		"status": party.IsOk,
	})
}

// POST /order/:partyID
func CreateClientOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	client := models.Client{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var item models.ItemOrder
	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err.Error())
	}

	item.Owner = client

	party := repository.GetPartyRepository().Get(partyId)
	if !party.IsOk {
		c.JSON(403, gin.H{
			"message": "Party is no ready",
			"party-tag": party.TAG,
			"party-status": party.IsOk,
		})
		return
	}
	party.AddClientOrder(item, client)

	c.JSON(201, gin.H{
		"message": "Order added",
		"party-tag": party.TAG,
		"Orders": party.GetClientOrder(client),
	})
}

// GET /order/:partyID
func GetClientOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	client := models.Client{}

	party := repository.GetPartyRepository().Get(partyId)

	c.JSON(200, gin.H{
		"message": "Order fetched",
		"party-tag": party.TAG,
		"Orders": party.GetClientOrder(client),
	})
}

// GET /party-order/:partyID
func GetAllOrder(c *gin.Context) {
	partyId := c.Param("partyID")

	party := repository.GetPartyRepository().Get(partyId)

	c.JSON(200, gin.H{
		"message": "Complete Order fetched",
		"party-tag": party.TAG,
		"Orders": party.GetCompleteOrder(),
	})
}

// POST /prepare-order/:partyID
func SendPrepareCommandOrder(c *gin.Context) {
	partyId := c.Param("partyID")

	party := repository.GetPartyRepository().Get(partyId)

	// TODO Validate sender is host?
	client := models.Client{}
	if party.GetHost().Id != client.Id {
		c.JSON(400, gin.H{
			"message": "Action must be performed by party host",
			"party-tag": party.TAG,
		})
		return
	}

	if !party.IsOk {
		c.JSON(403, gin.H{
			"message": "Party is no ready",
			"party-tag": party.TAG,
			"party-status": party.IsOk,
		})
		return
	}

	// TODO trigger goroutine to save orders to order-management-service

	c.JSON(201, gin.H{
		"message": "Orders pushed",
		"party-tag": party.TAG,
		"party": party,
	})
}
