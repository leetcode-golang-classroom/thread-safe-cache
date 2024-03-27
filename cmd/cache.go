package main

import (
	"log"
	"sync"
)

type Cache struct {
	sync.RWMutex
	data map[string]any
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]any),
	}
}

func (m *Cache) Get(key string) (any, bool) {
	m.RLock()
	defer m.RUnlock()
	val := m.data[key]
	return val, val != nil
}

func (m *Cache) Set(key string, value any) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func (m *Cache) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func (m *Cache) Contains(key string) bool {
	m.RLock()
	defer m.RUnlock()
	val := m.data[key]
	return val != nil
}

func (m *Cache) Keys() []string {
	m.RLock()
	defer m.RLock()
	keys := make([]string, 0)
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func RunCacheExample() {
	cache := NewCache()
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	keys := cache.Keys()
	for k := range keys {
		log.Printf("key: %v", k)
	}

	a, _ := cache.Get("a")
	log.Printf("a: %v", a)
	b, _ := cache.Get("b")
	log.Printf("b: %v", b)
	z, _ := cache.Get("z")
	log.Printf("z: %v", z)

	cache.Delete("a")
	cache.Delete("z")

	a, exists := cache.Get("a")
	log.Printf("a: %v, exists: %v", a, exists)
	for k := range keys {
		log.Printf("key: %v", k)
	}
}
