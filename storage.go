package main

import "fmt"

// StorageType is a enum to define available storages types
type StorageType int

const (
	// Nil is a null value to StorageType when dont have storage
	Nil StorageType = iota
	// JSON will store data in JSON files saved on disk
	JSON
	// Zabbix will store in Zabbix Server configured
	Zabbix
	// Memory will store the data in-memory cache
	Memory
)

// Storage represents all possible actions available to deal with data
type Storage interface {
	SaveMetric(Metric) error
	SaveItem(Metric) error
	ItemExists(Metric) bool
}

// NewStorageType return a StorageType based on string
func NewStorageType(s string) (stgType StorageType, err error) {
	switch s {
	case "JSON":
		stgType = JSON
	case "Zabbix":
		stgType = Zabbix
	default:
		err = fmt.Errorf("STORAGE_TYPE is not supported: %s", s)
	}
	return
}

// NewStorage returns a interface with the storage type choosed
func NewStorage(storageType StorageType, storageURL string) (Storage, error) {
	var storage Storage
	var err error

	switch storageType {
	case Memory:
		storage, err = NewStorageMemory()
	case JSON:
		storage, err = NewStorageJSON(storageURL)
	case Zabbix:
		storage, err = NewStorageZabbixSender(storageURL)
	default:
		err = fmt.Errorf("STORAGE_TYPE is not supported: %v", storageType)
	}

	return storage, err
}
