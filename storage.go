package main

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
func NewStorageType(s string) StorageType {
	switch s {
	case "JSON":
		return JSON
	case "Zabbix":
		return Zabbix
	// the default STORAGE_TYPE if not set
	default:
		return JSON
	}
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
		//storage, err = NewStorageZabbix()
	}

	return storage, err
}
