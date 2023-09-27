package main

import (
	"fmt"
	"sync"
)

var m map[string]int
var mu sync.Mutex

// Here, 5 goroutines are writing to a shared map data. As a result: unexpected
// output, sometimes = fatal error: concurrent map writes

func writeToMap(key string, value int) {
	//mu.Lock() // We can add Mutex to prevent race conditions. It ensures that only
	//one goroutine can access it at a time.
	m[key] = value
	//mu.Unlock()
}

func main() {
	m = make(map[string]int)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			writeToMap(fmt.Sprintf("key%d", n), n)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
