package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.RWMutex{},
	}

	ticker := time.NewTicker(interval)
 	go func(){
		for {
			select {
				case <-ticker.C:
					cache.reapLoop(interval)
			}
		}
	}()

	return cache 
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.cache[key]
	if !found {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for key, entry := range c.cache {
		if time.Now().Sub(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
