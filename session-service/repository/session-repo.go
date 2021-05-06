package repository

import (
	"errors"
	"klauskie.com/microservice-aurant/session-service/models"
)

var instance *repoSession = nil

type SessionRepository interface {
	Add(user *models.SimpleUser)
	Get(tag string) *models.SimpleUser
	Remove(tag string) error
	Update(party *models.SimpleUser) error
	GetAll() map[string]*models.SimpleUser
	ClearAll()
}

// SessionRepository
func GetSessionRepository() SessionRepository {
	if instance == nil {
		instance = new(repoSession)
		instance.collection = make(map[string]*models.SimpleUser)
	}
	return instance
}

// RepoModel
type repoSession struct {
	collection		map[string]*models.SimpleUser
}

func (r repoSession) Add(party *models.SimpleUser) {
	r.collection[party.Token] = party
}

func (r repoSession) Get(tag string) *models.SimpleUser {
	if _, ok := r.collection[tag]; ok {
		return r.collection[tag]
	}
	return nil
}

func (r repoSession) Remove(tag string) error {
	if _, ok := r.collection[tag]; ok {
		delete(r.collection, tag)
		return nil
	}
	return errors.New("could not find User")
}

func (r repoSession) Update(party *models.SimpleUser) error {
	if _, ok := r.collection[party.Token]; ok {
		r.collection[party.Token] = party
		return nil
	}
	return errors.New("could not find User, create user first")
}

func (r repoSession) GetAll() map[string]*models.SimpleUser {
	return r.collection
}

func (r repoSession) ClearAll() {
	for key, _ := range r.collection {
		delete(r.collection, key)
	}
	// r.collection = make(map[string]*models.Party)
}