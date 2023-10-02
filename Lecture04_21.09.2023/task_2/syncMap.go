package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := NewOwnSyncMap()
	mu.Store("Shyntas", 19)
	val, ok := mu.Load("Shyntas")
	fmt.Println(val, ok)

	mu.Delete("Shyntas")
	val, ok = mu.Load("Shyntas")
	fmt.Println(val, ok)
}

type OwnSyncMap struct {
	rwm sync.RWMutex
	mp  map[string]int
}

func NewOwnSyncMap() *OwnSyncMap {
	return &OwnSyncMap{
		mp: make(map[string]int),
	}
}

func (osp *OwnSyncMap) Load(key string) (int, bool) {
	osp.rwm.RLock()
	defer osp.rwm.RUnlock()
	val, ok := osp.mp[key]
	return val, ok
}

func (osp *OwnSyncMap) Store(key string, val int) {
	osp.rwm.Lock()
	defer osp.rwm.Unlock()
	osp.mp[key] = val
}

func (osp *OwnSyncMap) Delete(key string) {
	osp.rwm.Lock()
	defer osp.rwm.Unlock()
	delete(osp.mp, key)
}
