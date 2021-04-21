package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"klauskie.com/microservice-aurant/party-service/util"
)

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	respData := struct {
		User struct {
			Name string `json:"name"`
			VendorID string `json:"vendor_id"`
		} `json:"user"`
	}{}
	if err := json.Unmarshal(body, &respData); err != nil {
		fmt.Println(err.Error())
	}

	c.Set("clientName", respData.User.Name)
	c.Set("vendorID", respData.User.VendorID)
	c.Set("token", token)

	c.Next()
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
