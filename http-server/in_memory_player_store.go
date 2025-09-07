package main

import "sync"

type InMemoryPlayerStore struct {
	mc    sync.RWMutex
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{sync.RWMutex{}, map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mc.RLock()
	defer i.mc.RUnlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mc.Lock()
	defer i.mc.Unlock()
	i.store[name]++
}
