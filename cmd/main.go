package main

import "log"

func main() {
	cache := NewShardMap(3)
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
