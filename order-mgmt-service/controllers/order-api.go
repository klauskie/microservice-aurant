package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/microservice-aurant/order-mgmt-service/models"
	"klauskie.com/microservice-aurant/order-mgmt-service/repository"
	"log"
)


// POST /vendor/:restaurantId
func SaveOrderBatch(c *gin.Context) {

	log.Println("SaveOrderBatch called")

	// Read item batch
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var batch models.OrderBatch
	if err := json.Unmarshal(body, &batch); err != nil {
		fmt.Println(err.Error())
	}

	log.Println(batch)

	// save
	repository.GetOrderRepository().Add(&batch)

	// notify vendor
	c.JSON(201, gin.H{
		"service": "Order Management Service",
	})
}
