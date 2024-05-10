package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Locations map[string]cacheEntry
	Mutex     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Locations: make(map[string]cacheEntry),
		Mutex:     &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c

}

func (c Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	v := cacheEntry{createdAt: time.Now(), val: val}
	c.Locations[key] = v
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	val, contains := c.Locations[key]

	return val.val, contains
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c Cache) reap(now time.Time, interval time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for k, v := range c.Locations {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.Locations, k)
		}
	}
}
