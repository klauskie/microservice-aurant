package models

type SimpleUser struct {
	Token string `json:"token"`
	Name string `json:"name"`
	VendorID string `json:"vendor_id"`
	TimeStamp int64 `json:"time_stamp"`
}

func NewSimpleUser(token, name, vendorD string, createdAt int64) SimpleUser {
	return SimpleUser{
		Token:     token,
		Name:      name,
		VendorID:  vendorD,
		TimeStamp: createdAt,
	}
}