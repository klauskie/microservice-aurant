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

type ClientOrderDefinitionWrapper struct {
	Client Client `json:"client"`
	OrderList []CatalogItemOrder `json:"order_list"`
}

type CatalogItemOrder struct {
	CatalogItem CatalogItem `json:"catalog_item"`
	Metadata ItemOrder `json:"metadata"`
}

type CatalogItem struct {
	ItemId string `json:"itemId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price string `json:"price"`
	IsAvailable bool `json:"isAvailable"`
	IsDisplayable bool `json:"isDisplayable"`
}