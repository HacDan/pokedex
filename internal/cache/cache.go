package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Locations map[string]cacheEntry
	Mutex     *sync.Mutex
	interval  time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{interval: interval}
	c.reapLoop()

	return c

}

func (c Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	v := cacheEntry{createdAt: time.Now(), val: val}
	c.Locations[key] = v
	c.Mutex.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	val, contains := c.Locations[key]
	c.Mutex.Unlock()

	return val.val, contains
}

func (c Cache) reapLoop() {
	for k, v := range c.Locations {
		if time.Now-c.interval > v.createdAt {
			delete(c.Locations, k)
		}
	}
}
