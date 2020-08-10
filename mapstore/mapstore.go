package mapstore

import (
	"errors"
	"github.com/rashmi43/go-messenger/domain"
)

type MapStore struct {
	store map[string]domain.Message // An in-memory store with a map
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Message)}
}

func (m MapStore) Create(c domain.Message) error {
	messageID := c.ID
	if _, ok := m.store[messageID]; ok {
		return errors.New("The message id already exists in store")
	}
	m.store[messageID] = c
	return nil
}

func (m MapStore) Update(newid string, c domain.Message) error {
	messageID := newid
	if _, ok := m.store[messageID]; ok {
		return errors.New("The message with the given ID doesn't exist")
	}
	m.store[messageID] = c
	return nil
}

func (m MapStore) Delete(id string) error {
	messageID := id
	if _, ok := m.store[messageID]; !ok {
		return errors.New("The message with the given ID doesn't exist")
	}
	delete(m.store, messageID)
	return nil

}

func (m MapStore) GetAll() ([]domain.Message, error) {
	messageList := []domain.Message{}
	for _, v := range m.store {
		messageList = append(messageList, v)
	}
	return messageList, nil

}

func (m MapStore) GetById(id string) (domain.Message, error) {
	messageID := id
	if message, ok := m.store[messageID]; !ok {
		return domain.Message{}, errors.New("The message with the given ID doesn't exist")
	} else {
		return message, nil
	}

}
