package main

import "sync"

func main() {
	// example 1:
	// read on nil / open and empty channels causes deadlock
	var nilCh chan int
	<-nilCh

	opEmCh := make(chan int)
	<-opEmCh

	// example 2:
	// write on nil / open and full channels causes deadlock
	nilCh <- 5

	chFull := make(chan int, 1)
	chFull <- 42
	go func() {
		chFull <- 99 // This line will never execute
	}()

	// example 3:
	// blocking on mutex: holding a different mutex and waiting each other to release their locks
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		mu2.Lock() // Deadlock

		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		mu1.Lock() // Deadlock

		mu1.Unlock()
		mu2.Unlock()
	}()

	// example 4:
	// wait group deadlock
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			// ...
		}()
	}

	wg.Wait() // Deadlock if Add and Done are not balanced correctly
}
