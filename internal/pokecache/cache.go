package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	data     map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	cacheEnt := cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Lock()
	c.data[key] = cacheEnt
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	if value, isMapContainsKey := c.data[key]; isMapContainsKey {
		c.mu.Unlock()
		return value.val, true
	} else {
		c.mu.Unlock()
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	go func() {
		for range ticker.C {
			c.mu.Lock()

			for key, entry := range c.data {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	newcache := Cache{
		interval: interval,
		data:     make(map[string]cacheEntry),
	}
	(*Cache).reapLoop(&newcache)
	return &newcache
}
