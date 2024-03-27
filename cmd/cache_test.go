package main

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	// concurrency fail race test
	for i := 1; i <= 10; i++ {
		go func(val int) {
			cache.Set(fmt.Sprint(val), val)
		}(i)
	}
}
