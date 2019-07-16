package main

import (
	"encoding/json"

	"github.com/patrickmn/go-cache"
)

type StorageMemory struct {
	cache *cache.Cache
}

// NewStorageMemory create a new cache in-memory
func NewStorageMemory() (*StorageMemory, error) {
	var err error
	stg := new(StorageMemory)
	stg.cache = cache.New(cache.NoExpiration, cache.NoExpiration)

	return stg, err
}

// SaveMetric save the metric in-memory cache
func (s *StorageMemory) SaveMetric(metric Metric) error {
	return s.SaveItem(metric)
}

// SaveItem save the item in-memory cache
func (s *StorageMemory) SaveItem(metric Metric) error {
	var err error

	// transform to json
	jsonMetric, err := json.Marshal(metric)
	if err != nil {
		return err
	}

	s.cache.Set(metric.Stats.Name, string(jsonMetric), cache.NoExpiration)
	return err
}

// ItemExists check if a item exists in-memory cache
func (s *StorageMemory) ItemExists(metric Metric) (found bool) {
	_, found = s.cache.Get(metric.Stats.Name)
	return
}
