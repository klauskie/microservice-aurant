package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"klauskie.com/microservice-aurant/session-service/repository"
	"time"
)

func IsTokenValid(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(SECRET_SESSION_KEY), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		simpleUser := repository.GetSessionRepository().Get(tokenString)
		if simpleUser.TimeStamp + TWENTY_FOUR_HOURS < time.Now().UnixNano() {
			return true
		}
	}
	repository.GetSessionRepository().Remove(tokenString)
	return false
}
