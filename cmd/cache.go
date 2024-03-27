package main

import (
	"crypto/sha1"
	"log"
	"sync"
)

type Shard struct {
	sync.RWMutex
	data map[string]any
}

type ShardMap []*Shard

func NewShardMap(n int) ShardMap {
	shards := make([]*Shard, n)
	for i := 0; i < n; i++ {
		shards[i] = &Shard{
			data: make(map[string]any),
		}
	}

	return shards
}

func (m ShardMap) getShard(key string) *Shard {
	// find index
	i := m.getShardIndex(key)
	return m[i]
}
func (m ShardMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[0])
	i := hash % len(m)
	log.Printf("key: %v, index: %v", key, i)
	return i
}
func (m ShardMap) Get(key string) (any, bool) {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val := shard.data[key]
	return val, val != nil
}

func (m ShardMap) Set(key string, value any) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

func (m ShardMap) Delete(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

func (m ShardMap) Contains(key string) bool {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val := shard.data[key]
	return val != nil
}

func (m ShardMap) Keys() []string {
	keys := make([]string, 0)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(len(m))
	for _, shard := range m {
		go func(s *Shard) {
			s.RLock()
			for k := range s.data {
				mutex.Lock()
				keys = append(keys, k)
				mutex.Unlock()
			}
			s.RUnlock()
			wg.Done()
		}(shard)
	}
	return keys
}
