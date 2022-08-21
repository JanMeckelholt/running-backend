package persistence

import "github.com/JanMeckelholt/running/tree/main/running-backend/config"

type Storer interface {
	InitStorage() error
	AutoMigrate(interface{}) error
	Create(interface{}) error
	Get(string, interface{}) (interface{}, error)
	GetObjectById(key1, key2 string, id uint, obj interface{}) (interface{}, error)
	UpdateObject(existing, object interface{}) error
	DeleteObjectById(key uint, object interface{}) error
	AddObjectToForeignKey(id uint, object interface{}) error
	GetWhere(string, interface{}, interface{}) (interface{}, error)
}

func NewStorage(conf config.PersistenceConfig) Storer {
	return NewSQLStorer(conf)
}
