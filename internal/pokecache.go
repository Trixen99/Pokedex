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
	nCache := Cache{
		data:     map[string]cacheEntry{},
		interval: time,
	}
	return &nCache
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
