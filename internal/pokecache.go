package internal

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	interval time.Duration
	sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(time time.Duration) *Cache {
	nCache := &Cache{
		data:     map[string]cacheEntry{},
		interval: time,
	}
	go nCache.reapLoop()
	return nCache
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	valByte, ok := c.data[key]
	if ok {
		return valByte.val, true
	} else {
		return []byte{}, false
	}

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.Lock()
		for key, entry := range c.data {
			c.Lock()
			if time.Since(entry.createdAt) >= c.interval {
				delete(c.data, key)
			}

		}
		c.Unlock()
	}
}
