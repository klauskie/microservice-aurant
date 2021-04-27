package models

import (
	"bytes"
	"errors"
	"math/rand"
	"sort"
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

func (p *Party) AddClientOrder(item ItemOrder, clientId string) {
	clientOrder := p.OrderMap[clientId]
	clientOrder = append(clientOrder, item)
	p.OrderMap[clientId] = clientOrder
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

func (p *Party) GetCompleteOrder() []ClientOrderWrapper {
	var wrappers []ClientOrderWrapper
	for key, itemList := range p.OrderMap {
		client, clientError := p.GetClientByID(key)
		if clientError != nil {
			continue
		}
		wrap := ClientOrderWrapper{
			Client:    client,
			OrderList: itemList,
		}
		wrappers = append(wrappers, wrap)
	}
	return wrappers
}

func (p *Party) GetCompleteOrder_withDefinition(itemCatalogMap map[string]CatalogItem, clientId string) []ClientOrderDefinitionWrapper {
	var wrappers []ClientOrderDefinitionWrapper
	for key, itemList := range p.OrderMap {
		client, clientError := p.GetClientByID(key)
		if clientError != nil {
			continue
		}

		var itemWraps []CatalogItemOrder
		for _, item := range itemList {

			itemWrap := CatalogItemOrder{
				CatalogItem: itemCatalogMap[item.ItemId],
				Metadata:    item,
			}
			itemWraps = append(itemWraps, itemWrap)
		}
		wrap := ClientOrderDefinitionWrapper{
			Client:    client,
			OrderList: itemWraps,
		}
		wrappers = append(wrappers, wrap)
	}

	sort.Slice(wrappers, func(i, j int) bool {
		return wrappers[i].Client.Id == clientId
	})

	return wrappers
}

func (p *Party) GetClientOrder_withDefinition(clientId string, itemCatalogMap map[string]CatalogItem) (ClientOrderDefinitionWrapper, error) {
	client, clientError := p.GetClientByID(clientId)
	if clientError != nil {
		return ClientOrderDefinitionWrapper{}, clientError
	}

	var itemWraps []CatalogItemOrder
	for _, item := range p.OrderMap[clientId] {

		itemWrap := CatalogItemOrder{
			CatalogItem: itemCatalogMap[item.ItemId],
			Metadata:    item,
		}
		itemWraps = append(itemWraps, itemWrap)
	}
	wrap := ClientOrderDefinitionWrapper{
		Client:    client,
		OrderList: itemWraps,
	}
	return wrap, nil
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

func (p *Party) GetClientByID(clientID string) (Client, error) {
	// Verify content instead of key
	list := p.GetClients()
	for _, client := range list {
		if client.Id == clientID {
			return client, nil
		}
	}
	return Client{}, errors.New("No Client found with given ID")
}

func (p *Party) GetAllItemIds() []string {
	var itemIds []string
	for _, items := range p.OrderMap {
		for _, item := range items {
			itemIds = append(itemIds, item.ItemId)
		}
	}
	return itemIds
}

func (p *Party) GetClientItemIds(clientId string) []string {
	var itemIds []string

	for _, item := range p.OrderMap[clientId] {
		itemIds = append(itemIds, item.ItemId)
	}

	return itemIds
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
