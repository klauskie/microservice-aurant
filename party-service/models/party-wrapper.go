package models

type PartyWrapper struct {
	TAG string `json:"tag"`
	ClientList []Client `json:"client_list"`
	RestaurantId string `json:"restaurant_id"`
	Host Client `json:"host"`
	OrderMap map[string][]ItemOrder `json:"order_map"`
	IsOk bool `json:"is_ok"`
}

func NewPartyWrapper(party Party) PartyWrapper {
	return PartyWrapper{
		TAG:          party.TAG,
		ClientList:   party.GetClients(),
		RestaurantId: party.RestaurantId,
		Host:         party.GetHost(),
		OrderMap:     party.OrderMap,
		IsOk:         party.IsOk,
	}
}