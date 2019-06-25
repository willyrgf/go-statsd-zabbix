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
	SaveMetric(Metric) error
}

// NewStorage returns a interface with the storage type choosed
func NewStorage(storageType StorageType, storageURL string) (Storage, error) {
	var storage Storage
	var err error

	switch storageType {
	case JSON:
		storage, err = NewStorageJSON(storageURL)
	case Zabbix:
		//storage, err = NewStorageZabbix()
	}

	return storage, err
}
