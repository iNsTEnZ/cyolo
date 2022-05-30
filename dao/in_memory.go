package dao

import "sync"

type InMemory struct {
	storage sync.Map
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (dao *InMemory) Put(key interface{}, value interface{}) error {
	dao.storage.Store(key, value)
	return nil
}

func (dao *InMemory) Get(key interface{}, defaultValue interface{}) (interface{}, error) {
	if value, exists := dao.storage.Load(key); exists {
		return value, nil
	}

	return defaultValue, nil
}
