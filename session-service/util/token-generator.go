package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"klauskie.com/microservice-aurant/session-service/repository"
	"time"
)

func IsTokenValid(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(SECRET_SESSION_KEY), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		simpleUser := repository.GetSessionRepository().Get(tokenString)
		if simpleUser == nil {
			return false
		}

		delta := time.Now().Unix() - simpleUser.TimeStamp
		if delta < TWENTY_FOUR_HOURS {
			return true
		} else {
			repository.GetSessionRepository().Remove(tokenString)
		}
	}
	return false
}
