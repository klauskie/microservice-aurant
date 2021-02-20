package models

type ItemOrder struct {
	ItemId string `json:"item_id"`
	Instructions string `json:"instructions"`
	Quantity int `json:"quantity"`
	Owner Client
}
