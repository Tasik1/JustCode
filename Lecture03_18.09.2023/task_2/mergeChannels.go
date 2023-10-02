package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			a <- 1
		}
		close(a)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			b <- 2
		}
		close(b)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			c <- 3
		}
		close(c)
	}()
	for v := range mergeChannels(a, b, c) {
		fmt.Println(v)
	}
}

func mergeChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch <-chan int) {
			for val := range ch {
				mergedCh <- val
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func isSlicesEquals2(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	counts := make(map[int]int)
	for _, val := range arr1 {
		counts[val]++
	}

	for _, val := range arr2 {
		if count, ok := counts[val]; !ok || count == 0 {
			return false
		}
		counts[val]--
	}

	return true
}
