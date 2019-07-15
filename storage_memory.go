package main

import (
	"github.com/patrickmn/go-cache"
)

type StorageMemory struct {
	cache *cache.Cache
}
