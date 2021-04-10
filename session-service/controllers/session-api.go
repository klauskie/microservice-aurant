package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/microservice-aurant/session-service/models"
	"klauskie.com/microservice-aurant/session-service/repository"
	"klauskie.com/microservice-aurant/session-service/util"
	"time"
)

// TODO Register account
// POST /register
func RegisterClient(c *gin.Context) {

}

// TODO Sign in with account
// POST /login
func Login(c *gin.Context) {

}

// TODO sign in as guest
// POST /guest-login
func LoginAsGuest(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	userData := struct {
		Name string `json:"name"`
		VendorID string `json:"vendor_id"`
	}{}
	if err := json.Unmarshal(body, &userData); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Invalid post data",
		})
		return
	}

	createdAt := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"vendor_id":  userData.VendorID,
		"name": userData.Name,
		"created_at":  createdAt,
	})

	tokenString, err := token.SignedString([]byte(util.SECRET_SESSION_KEY))

	simpleUser := models.NewSimpleUser(tokenString, userData.Name, userData.VendorID, createdAt)
	repository.GetSessionRepository().Add(&simpleUser)

	c.JSON(202, gin.H{
		"message": "Guest user registered",
		"token": tokenString,
		"user": userData,
	})
}

// DELETE /logout
func Logout(c *gin.Context) {
	tokenString := c.Request.Header.Get("token")
	repository.GetSessionRepository().Remove(tokenString)

	c.JSON(202, gin.H{
		"message": "User logged out",
	})
}

// GET /profile
func Profile(c *gin.Context) {
	tokenString := c.Request.Header.Get("token")
	if !util.IsTokenValid(tokenString) {
		c.JSON(401, gin.H{
			"message": "Invalid token",
		})
		return
	}
	user := repository.GetSessionRepository().Get(tokenString)
	c.JSON(202, gin.H{
		"message": "Retrieved user",
		"user": user,
	})
}

// GET /token-validation/:token
func TokenValidation(c *gin.Context) {
	tokenString := c.Param("token")
	if !util.IsTokenValid(tokenString) {
		c.JSON(401, gin.H{
			"message": "Invalid token",
		})
		return
	}
	user := repository.GetSessionRepository().Get(tokenString)
	c.JSON(202, gin.H{
		"message": "Retrieved user",
		"user": user,
	})
}
