package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	createdtime := time.Now()
	newCacheEntry := cacheEntry{createdAt: createdtime, val: val}
	c.entries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}

}
