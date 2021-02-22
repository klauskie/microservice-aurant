package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/order-mgmt-service/controllers"
)

func main() {
	r := gin.Default()

	api := r.Group("/api-order")
	{
		api.POST("/batch", controllers.SaveOrderBatch)
	}

	r.Run(":8082")
}
