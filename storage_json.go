package main

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// Collection identifier for JSON collection about Metric
	Collection = "metrics"
	// DefaultPathStorageJSON is a default path to storage JSON
	// when the STORAGE_URL env/config is not set
	DefaultPathStorageJSON = "./data/"
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

func getFileNameJSON(metric Metric) string {
	timestampFormat := "20060102_150405.000"
	timeOfMetric := metric.Timestamp.Format(timestampFormat)
	name := metric.Prefix
	if len(name) < 1 {
		name = metric.Hostname
	}
	return fmt.Sprintf("%s_%s", name, timeOfMetric)
}

// SaveMetric save a new metric
func (s *StorageJSON) SaveMetric(metric Metric) error {
	filename := getFileNameJSON(metric)
	if err := s.db.Write(Collection, filename, metric); err != nil {
		return err
	}

	return nil
}
