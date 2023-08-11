package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, ok := c.cache[key]

	if !ok {
		return []byte{}, ok
	} else {
		return data.val, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTimer(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func NewCache(d time.Duration) Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}

	go cache.reapLoop(d)

	return cache
}
