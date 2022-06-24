package main

import (
	"fmt"
	"sync"
)

// Реализовать паттерн «адаптер» на любом примере.

type Cache interface {
	Get(key string) (string, bool)
	Set(key string, val string)
}

type CacheMap struct {
	sync.RWMutex
	m map[string]string
}

func (c *CacheMap) Set(key string, val string) {
	c.Lock()
	c.m[key] = val
	c.Unlock()
}

func (c *CacheMap) Get(key string) (string, bool) {
	c.RLock()
	val, ok := c.m[key]
	c.RUnlock()
	return val, ok
}

type CacheSyncMap struct {
	m sync.Map
}

func (c *CacheSyncMap) Set(key string, val string) {
	c.m.Store(key, val)
}

func (c *CacheSyncMap) Get(key string) (string, bool) {
	val, ok := c.m.Load(key)
	if !ok {
		return "", ok
	}
	return val.(string), ok
}

func do(c Cache) {
	key := "key"
	val := "Hello World"
	c.Set(key, val)
	v, _ := c.Get(key)
	fmt.Println(v)
}

func main() {
	c1 := &CacheMap{m: make(map[string]string)}
	c2 := &CacheSyncMap{}
	do(c1)
	do(c2)
}
