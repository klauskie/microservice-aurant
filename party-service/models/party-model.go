package models

import (
	"bytes"
	"math/rand"
	"time"
)

type Party struct {
	TAG string
	RestaurantId string
	ClientSet map[string]Client
	OrderMap map[string][]ItemOrder
	IsOk bool
}

func NewParty(vendorUUID string, host Client) *Party {
	p := &Party{
		TAG:          createTag(),
		RestaurantId: vendorUUID,
		ClientSet:   make(map[string]Client),
		OrderMap:     make(map[string][]ItemOrder),
		IsOk: false,
	}
	p.ClientSet["host"] = host
	return p
}

func (p *Party) AddClient(client Client) {
	if _, ok := p.ClientSet[client.Id]; !ok {
		p.OrderMap[client.Id] = []ItemOrder{}
		p.ClientSet[client.Id] = client
	}
}

func (p *Party) AddClientOrder(item ItemOrder, client Client) {
	clientOrder := p.OrderMap[client.Id]
	clientOrder = append(clientOrder, item)
	p.OrderMap[client.Id] = clientOrder
}

func (p *Party) addClientOrders(items []ItemOrder, client Client) {
	clientOrder := p.OrderMap[client.Id]
	clientOrder = append(clientOrder, items...)
	p.OrderMap[client.Id] = clientOrder
}

func (p *Party) RemoveClient(clientID string) {
	items := p.OrderMap[clientID]
	delete(p.ClientSet, clientID)
	p.addClientOrders(items, p.GetHost())
}

func (p *Party) GetCompleteOrder() []ItemOrder {
	var items []ItemOrder
	for _, itemList := range p.OrderMap {
		items = append(items, itemList...)
	}
	return items
}

func (p *Party) GetClientOrder(client Client) []ItemOrder {
	return p.OrderMap[client.Id]
}

func (p *Party) GetHost() Client {
	return p.ClientSet["host"]
}

func (p *Party) GetClients() []Client {
	var list []Client
	for _, client := range p.ClientSet {
		list = append(list, client)
	}
	return list
}

func (p *Party) IsClientOnParty(clientID string) bool {
	// Verify content instead of key
	list := p.GetClients()
	for _, client := range list {
		if client.Id == clientID {
			return true
		}
	}
	return false
}

// TODO Avoid collisions
func createTag() string {
	rand.Seed(time.Now().UnixNano())
	buff := bytes.NewBufferString("")
	for i := 0; i < 4; i++ {
		letter := rand.Intn(90 - 65) + 65
		buff.WriteString(string(rune(letter)))
	}
	return buff.String()
}
