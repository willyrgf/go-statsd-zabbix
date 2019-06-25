package main

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// Collection identifier for JSON collection about Metric
	Collection = "metrics"
	// DefaultPathStorageJSON is a default path to storage JSON
	// when the STORAGE_URL env/config is not set
	DefaultPathStorageJSON = "~/tmp/"
)

// StorageJSON is the data storage layer using JSON file
type StorageJSON struct {
	db *scribble.Driver
}

// NewStorageJSON create a new driver to storage JSON in file
func NewStorageJSON(location string) (*StorageJSON, error) {
	var err error

	stg := new(StorageJSON)

	if len(location) < 1 {
		location = DefaultPathStorageJSON
	}

	stg.db, err = scribble.New(location, nil)
	if err != nil {
		return nil, err
	}

	return stg, nil
}

// SaveMetric save a new metric
func (s *StorageJSON) SaveMetric(metric Metric) error {
	if err := s.db.Write(Collection, "ID10", metric); err != nil {
		return err
	}

	return nil
}
