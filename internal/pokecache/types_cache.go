package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}
