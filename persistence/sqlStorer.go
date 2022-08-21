package persistence

import (
	"errors"
	"github.com/JanMeckelholt/running/tree/main/running-backend/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLStorer struct {
	ConnString string
}

var DB *gorm.DB

func NewSQLStorer(persistenceConfig config.PersistenceConfig) Storer {
	store := SQLStorer{
		ConnString: persistenceConfig.Address,
	}
	err := store.InitStorage()
	if err != nil {
		panic(err.Error())
	}
	return store
}

func (ss SQLStorer) InitStorage() error {
	log.Info("Using ConnString for DB:", ss.ConnString)
	var err error
	db, err := gorm.Open(postgres.Open(ss.ConnString), &gorm.Config{})
	if err != nil {
		log.Error("DB error: could not open sql connection")
		panic("failed to connect to database")
	}
	DB = db
	log.Info("InitStorage finished")
	return nil
}

func (ss SQLStorer) AutoMigrate(object interface{}) error {
	log.Info("Starting automatic migrations")

	var err error
	if err = DB.Debug().AutoMigrate(object); err != nil {
		log.Error("DB error: could not automigrate object")
		return err
	}
	log.Info("Automatic migrations finished")
	return nil
}

func (ss SQLStorer) Create(object interface{}) error {
	result := DB.Create(object)
	if result.Error != nil {
		log.Error("DB error: could not create object")
		return result.Error
	}
	log.Tracef("Stored object")
	return nil
}

func (ss SQLStorer) Get(key string, object interface{}) (interface{}, error) {
	var result *gorm.DB
	if key == "" {
		result = DB.Find(object)
	} else {
		result = DB.Preload(key).Find(object)
	}
	if result.Error != nil {
		log.Error("DB error: could not get object")
		return nil, result.Error
	}
	log.Tracef("Retrieved object")
	return object, nil
}

func (ss SQLStorer) GetWhere(key string, clause interface{}, object interface{}) (interface{}, error) {
	var result *gorm.DB
	if key == "" {
		result = DB.Where(clause).Find(object)
	} else {
		log.Infof("whereClause in SQLStorer: %v", clause)
		result = DB.Preload(key).Where(clause).Find(object)
	}
	if result.Error != nil {
		log.Error("DB error: could not get object")
		return nil, result.Error
	}
	log.Tracef("Retrieved object")
	return object, nil
}

func (ss SQLStorer) GetObjectById(key1, key2 string, id uint, object interface{}) (interface{}, error) {
	var result *gorm.DB
	if key1 == "" {
		result = DB.First(object, id)
	} else if key2 == "" {
		result = DB.Preload(key1).First(object, id)
	} else {
		result = DB.Preload(key1).Preload(key2).First(object, id)
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Infof("No record found with id %v", id)
		return nil, nil
	}
	if result.Error != nil {
		log.Errorf("DB error: could not get object by id %d", id)
		return nil, result.Error
	}
	log.Tracef("Retrieved object by id")
	return object, nil
}

func (ss SQLStorer) UpdateObject(existing, object interface{}) error {
	result := DB.Model(existing).Updates(object)
	if result.Error != nil {
		log.Error("DB error: could not update object")
		return result.Error
	}
	log.Tracef("Updated object")
	return nil
}

func (ss SQLStorer) AddObjectToForeignKey(key uint, object interface{}) error {
	err := ss.Create(object)
	if err != nil {
		log.Error("DB error: could not add object to foreign key")
		return err
	}
	log.Tracef("Added object to foreign key")
	return nil
}

func (ss SQLStorer) DeleteObjectById(id uint, object interface{}) error {
	result := DB.Delete(object, id)
	if result.Error != nil {
		log.Error("DB error: could not delete object")
		return result.Error
	}
	log.Tracef("Deleted object: %v", object)

	return nil
}
