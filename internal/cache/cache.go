package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Locations map[string]cacheEntry
	Mutux     sync.Mutex
	interval  time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []bytes
}

func NewCache(interval time.Duration) Cache {
	c := Cache{interval: interval}
	c.reapLoop()

	return c

}

func (c Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	c[key] = val
	c.Mutex.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	val, contains := c[key]
	c.Mutex.Unlock()

	return val, contains
}

func (c Cache) reapLoop() {
	for k, v := range c {
		if time.Now-interval > v.createdAt {
			delete(c, k)
		}
	}
}
