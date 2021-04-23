package models

type ItemOrder struct {
	ItemId string `json:"item_id"`
	Instructions string `json:"instructions"`
	Quantity int `json:"quantity"`
	Owner Client
}

type ClientOrderWrapper struct {
	Client Client `json:"client"`
	OrderList []ItemOrder `json:"order_list"`
}
