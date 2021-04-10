package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/util"
)

func TokenAuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("token")

	if token == "" {
		respondWithError(c, 401, "API token required")
		return
	}

	resp, err := http.Get(util.URL_SESSION_API + "/token-validation/" + token)
	if err != nil {
		respondWithError(c, 401, "Cannot reach Session Service")
		return
	}

	if resp.StatusCode != http.StatusAccepted {
		respondWithError(c, 401, "Invalid API token")
		return
	}

	c.Next()
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
