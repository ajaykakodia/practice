package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	v, ok := c.m[k]
	c.mu.RUnlock()
	return v, ok
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	if c.m == nil {
		c.m = make(map[string]string)
	}
	c.m[k] = v
	c.mu.Unlock()
}

func main() {
	cache := Cache{}

	cache.Set("Ram", "Yadav")
	cache.Set("Raja", "Som")
	v, ok := cache.Get("Ram")
	if !ok {
		fmt.Println("Value not stored in cache.")
	}
	fmt.Println("Ram:", v)
}
