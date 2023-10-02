package main

import (
	"fmt"
	"sync"
)

var (
	mp  map[string]string
	rwm sync.RWMutex
)

func main() {
	mp = make(map[string]string)
	go writeData("key1", "value1")
	go readData("key1")
	go readData("key2")
	go readData("key1")
	go writeData("key2", "value2")
	go readData("key2")
	go readData("key1")
	go readData("key2")

	//time.Sleep(time.Second * 5)
	block := make(chan int)
	<-block
}

func readData(key string) {
	rwm.RLock() // Acquire a read lock (multiple readers can hold this)
	defer rwm.RUnlock()

	value := mp[key]
	fmt.Printf("Reading key: %s, Value: %s\n", key, value)
}

func writeData(key, value string) {
	rwm.Lock() // Acquire a write lock (exclusive access)
	defer rwm.Unlock()

	mp[key] = value
	fmt.Printf("Writing key: %s, Value: %s\n", key, value)
}
