package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
