package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	mergedCh := mergeChannels(ch1, ch2, ch3)

	expected := []int{1, 2, 3, 4, 5, 6}
	var actual []int
	for v := range mergedCh {
		actual = append(actual, v)
	}

	if !isSlicesEquals2(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
