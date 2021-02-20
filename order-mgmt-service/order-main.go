package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", HomePage)

	r.Run(":8082")
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "Order Management Service",
	})
}
