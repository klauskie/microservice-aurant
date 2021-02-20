package repository

import (
	"errors"
	"klauskie.com/microservice-aurant/party-service/models"
)

var instance *repoParty = nil

type PartyRepository interface {
	Add(party *models.Party)
	Get(tag string) *models.Party
	Remove(tag string) error
	Update(party *models.Party) error
	GetAll() map[string]*models.Party
}

// PartyRepository
func GetPartyRepository() PartyRepository {
	if instance == nil {
		instance = new(repoParty)
		instance.collection = make(map[string]*models.Party)
	}
	return instance
}

// RepoModel
type repoParty struct {
	collection		map[string]*models.Party
}

func (r repoParty) Add(party *models.Party) {
	r.collection[party.TAG] = party
}

func (r repoParty) Get(tag string) *models.Party {
	if _, ok := r.collection[tag]; ok {
		return r.collection[tag]
	}
	return nil
}

func (r repoParty) Remove(tag string) error {
	if _, ok := r.collection[tag]; ok {
		delete(r.collection, tag)
		return nil
	}
	return errors.New("could not find Party")
}

func (r repoParty) Update(party *models.Party) error {
	if _, ok := r.collection[party.TAG]; ok {
		r.collection[party.TAG] = party
		return nil
	}
	return errors.New("could not find Party, create party first")
}

func (r repoParty) GetAll() map[string]*models.Party {
	return r.collection
}
