package main

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var smallSlice intSlice
	t.Run("Small slices", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			t.Run(strconv.Itoa(i+1), func(t *testing.T) {
				slice := generateRandomSlice(10)
				smallSlice = slice
				result := make([]int, len(slice))
				copy(result, slice)
				smallSlice.quickSort(smallSlice, 0, len(smallSlice)-1)
				sort.Ints(result)
				if !isArrayEquals(smallSlice, result) {
					t.Errorf("expected result: %v != %v - real result", result, smallSlice)
				}
			})
		}
	})
	t.Run("Medium slices", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			t.Run(strconv.Itoa(i+1), func(t *testing.T) {
				slice := generateRandomSlice(100)
				smallSlice = slice
				result := make([]int, len(slice))
				copy(result, slice)
				smallSlice.quickSort(smallSlice, 0, len(smallSlice)-1)
				sort.Ints(result)
				if !isArrayEquals(smallSlice, result) {
					t.Errorf("expected result: %v != %v - real result", result, smallSlice)
				}
			})
		}
	})
	t.Run("Big slices", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			t.Run(strconv.Itoa(i+1), func(t *testing.T) {
				slice := generateRandomSlice(1000)
				smallSlice = slice
				result := make([]int, len(slice))
				copy(result, slice)
				smallSlice.quickSort(smallSlice, 0, len(smallSlice)-1)
				sort.Ints(result)
				if !isArrayEquals(smallSlice, result) {
					t.Errorf("expected result: %v != %v - real result", result, smallSlice)
				}
			})
		}
	})
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) // Generate random integers between 0 and 999
	}
	return slice
}
