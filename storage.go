package main

// StorageType is a enum to define available storages types
type StorageType int

const (
	// JSON will store data in JSON files saved on disk
	JSON StorageType = 0
	// Zabbix will store in Zabbix Server configured
	Zabbix StorageType = 1
)

// Storage represents all possible actions available to deal with data
type Storage interface {
}

// NewStorage returns a interface with the storage type choosed
func NewStorage(storageType StorageType) (Storage, error) {
	var storage Storage
	var err error

	switch storageType {
	case JSON:
		//storage, error = NewStorageJSON()
	case Zabbix:
		//storage, error = NewStorageZabbix()
	}

	return storage, err
}
