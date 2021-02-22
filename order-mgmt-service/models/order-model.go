package models

type OrderBatch struct {
	TAG string `json:"tag"`
	RestaurantId string `json:"restaurant_id"`
	Orders []ItemOrder `json:"orders"`
}

type ItemOrder struct {
	ItemId string `json:"item_id"`
	Instructions string `json:"instructions"`
	Quantity int `json:"quantity"`
}
