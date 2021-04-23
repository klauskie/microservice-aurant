package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/models"
	"klauskie.com/microservice-aurant/party-service/repository"
	"klauskie.com/microservice-aurant/party-service/util"
)

// GET /party/:vendorID
func CreateParty(c *gin.Context) {
	//vendorUUID := c.Param("vendorID")
	vendorUUID := c.GetString("vendorID") // From middleware
	token := c.GetString("token") // From middleware
	clientName := c.GetString("clientName") // From middleware

	party := models.NewParty(vendorUUID, models.NewClient(token, clientName))
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}
	repository.GetPartyRepository().Add(party)

	c.JSON(201, gin.H{
		"message": "Party created successfully",
		"party-tag": party.TAG,
		"party": models.NewPartyWrapper(*party),
	})
}

// GET /party/:partyID
func GetParty(c *gin.Context) {
	partyId := c.Param("partyID")
	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	pWrapper := models.NewPartyWrapper(*party)

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"party": pWrapper,
	})
}

// PUT /party/:partyID
func JoinParty(c *gin.Context) {
	partyId := c.Param("partyID")
	token := c.GetString("token") // From middleware
	clientName := c.GetString("clientName") // From middleware

	client := models.NewClient(token, clientName)

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	party.AddClient(client)

	c.JSON(200, gin.H{
		"message": "Joined to party successfully",
		"party-tag": party.TAG,
		"party": models.NewPartyWrapper(*party),
	})
}

// DELETE /party/:partyID/kick/:client
func KickFromParty(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.Param("clientID")

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
		})
		return
	}
	party.RemoveClient(clientID)

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
	clientID := c.GetString("token")

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"clients": party.GetClients(),
	})
}

// GET /party-status/:partyID
func GetPartyStatus(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "GET party",
		"party-tag": party.TAG,
		"status": party.IsOk,
	})
}

// PUT /party-status/:partyID
func UpdatePartyStatus(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	status := struct {
		Ready bool `json:"ready"`
	}{}
	if err := json.Unmarshal(body, &status); err != nil {
		fmt.Println(err.Error())
	}

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	// Validate sender is host
	if party.GetHost().Id != clientID {
		c.JSON(401, gin.H{
			"message": "Action must be performed by party host",
			"party-tag": party.TAG,
		})
		return
	}

	party.IsOk = status.Ready

	c.JSON(202, gin.H{
		"message": "Party status updated",
		"party-tag": party.TAG,
		"status": party.IsOk,
	})
}

// POST /order/:partyID
func CreateClientOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")
	clientName := c.GetString("clientName")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var item models.ItemOrder
	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err.Error())
	}

	tempClient := models.NewClient(clientID, clientName)
	item.Owner = tempClient

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
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
	party.AddClientOrder(item, tempClient)

	c.JSON(201, gin.H{
		"message": "Order added",
		"party-tag": party.TAG,
		"Orders": party.GetClientOrder(tempClient),
	})
}

// GET /order/:partyID
func GetClientOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")
	clientName := c.GetString("clientName")

	tempClient := models.NewClient(clientID, clientName)

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order fetched",
		"party-tag": party.TAG,
		"Orders": party.GetClientOrder(tempClient),
	})
}

// GET /party-order/:partyID
func GetAllOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if !party.IsClientOnParty(clientID) {
		c.JSON(401, gin.H{
			"message": "Given client does not belong in party",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Complete Order fetched",
		"party-tag": party.TAG,
		"Orders": party.GetCompleteOrder(),
	})
}

// POST /prepare-order/:partyID
func SendPrepareCommandOrder(c *gin.Context) {
	partyId := c.Param("partyID")
	clientID := c.GetString("token")

	party := repository.GetPartyRepository().Get(partyId)
	if party == nil {
		c.JSON(400, gin.H{
			"message": "No party found with given TAG",
		})
		return
	}

	if party.GetHost().Id != clientID {
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

	// HTTP post batch to api-order
	go sendOrderBatch(*party)

	c.JSON(201, gin.H{
		"message": "Orders pushed",
		"party-tag": party.TAG,
		"party": models.NewPartyWrapper(*party),
	})
}

// POST /test/prepare-order
func SendPrepareCommandOrderTest(c *gin.Context) {
	order1 := models.ItemOrder{
		ItemId:       "111",
		Instructions: "No onions",
		Quantity:     1,
	}

	order2 := models.ItemOrder{
		ItemId:       "222",
		Instructions: "",
		Quantity:     2,
	}

	order3 := models.ItemOrder{
		ItemId:       "333",
		Instructions: "extra napkins",
		Quantity:     1,
	}

	oMap := map[string][]models.ItemOrder{}
	oMap["alice"] = []models.ItemOrder{order1, order2}
	oMap["bob"] = []models.ItemOrder{order3}

	party := models.Party{
		TAG:          "TEST",
		RestaurantId: "1234",
		ClientSet:    nil,
		OrderMap:     oMap,
		IsOk:         true,
	}

	go sendOrderBatch(party)

	c.JSON(201, gin.H{
		"message": "Orders pushed",
		"party-tag": party.TAG,
		"party": models.NewPartyWrapper(party),
	})
}

func sendOrderBatch(party models.Party) {

	batch := struct{
		TAG string `json:"tag"`
		RestaurantId string `json:"restaurant_id"`
		Orders []models.ClientOrderWrapper `json:"orders"`
	}{
		TAG: party.TAG,
		RestaurantId: party.RestaurantId,
		Orders: party.GetCompleteOrder(),
	}

	requestBody, err := json.Marshal(batch)
	if err != nil {
		log.Fatal(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", util.URL_ORDER_API_SEND_BATCH, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
