package repository

import (
	"errors"
	"klauskie.com/microservice-aurant/order-mgmt-service/models"
)

var instance *repo = nil

type OrderRepository interface {
	Add(batch *models.OrderBatch)
	Get(tag string) *models.OrderBatch
	Update(party *models.OrderBatch) error
	GetAll(vendorID string) map[string]*models.OrderBatch
}

// OrderRepository
func GetOrderRepository() OrderRepository {
	if instance == nil {
		instance = new(repo)
		instance.collection = make(map[string]*models.OrderBatch)
	}
	return instance
}

// RepoModel
type repo struct {
	collection		map[string]*models.OrderBatch
}

func (r repo) Add(party *models.OrderBatch) {
	// TODO save batch to db
}

func (r repo) Get(tag string) *models.OrderBatch {
	// TODO get batch
	return nil
}

func (r repo) Update(party *models.OrderBatch) error {
	// TODO update batch
	return errors.New("could not find Order, create order first")
}

func (r repo) GetAll(restaurantID string) map[string]*models.OrderBatch {
	// TODO get batches by restaurant
	return r.collection
}

