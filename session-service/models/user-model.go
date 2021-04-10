package models

import (
	"time"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type SimpleUser struct {
	Token, Name, PartyID, VendorID string
	TimeStamp int64
}

func NewSimpleUser(token, name, partyID, vendorD string) SimpleUser {
	return SimpleUser{
		Token:     token,
		Name:      name,
		PartyID:   partyID,
		VendorID:  vendorD,
		TimeStamp: time.Now().UnixNano(),
	}
}