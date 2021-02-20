package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", HomePage)

	r.Run(":8081")
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "Party Service",
	})
}
