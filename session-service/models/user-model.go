package models

type SimpleUser struct {
	Token, Name, VendorID string
	TimeStamp int64
}

func NewSimpleUser(token, name, vendorD string, createdAt int64) SimpleUser {
	return SimpleUser{
		Token:     token,
		Name:      name,
		VendorID:  vendorD,
		TimeStamp: createdAt,
	}
}