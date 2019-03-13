package cache

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("key not found")

type Cache struct {
	mu     *sync.RWMutex
	cache  map[string][]byte
	prefix string
}

func NewCache(prefix string) *Cache {
	return &Cache{
		mu:     &sync.RWMutex{},
		cache:  make(map[string][]byte),
		prefix: prefix,
	}
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[c.prefix+key] = val
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if res, ok := c.cache[c.prefix+key]; ok {
		return res, nil
	}
	return nil, ErrNotFound
}
