package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	createdtime := time.Now()
	newCacheEntry := cacheEntry{createdAt: createdtime, Val: val}
	c.Entries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}

}
