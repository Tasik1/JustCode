package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := intSlice{5, 1, 3, 9, 6, 2, 7, -10, 0, 35, 46, -5, -4926, 5621, 76}
	fmt.Println("Unsorted array:", arr)

	arr.quickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}

// 1. https://leetcode.com/problems/two-sum/
// Solution: https://leetcode.com/submissions/detail/1048054530/
// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		val, ok := mp[target-num]
		if ok {
			return []int{val, i}
		}
		mp[num] = i
	}
	return nil
}

// 2. https://leetcode.com/problems/longest-common-prefix/
// Solution: https://leetcode.com/submissions/detail/1048130154/
// Time Complexity: O(n*m)
// Space Complexity: O(m)
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		for !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// 3. Написать функцию для сравнения слайса с цело-числовыми значениями
// a) сравнивает два слайса
// b) возвращает булево значение: совпало или нет
// c) порядок важен:
// Time Complexity: O(n)
// Space Complexity: O(1)
func isArrayEquals(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

// d) если порядок не важен:
// Time Complexity: O(n)
// Space Complexity: O(n)
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

// 4. Реализовать структуру с функцией для сортировки элементов слайса целых чисел
// Time Complexity: Average Case: O(n log n)
// Space Complexity: O(log n)
type intSlice []int

func (is intSlice) quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		is.quickSort(arr, low, pi-1)
		is.quickSort(arr, pi, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low+(high-low)/2]

	for low <= high {
		for arr[low] < pivot {
			low++
		}
		for arr[high] > pivot {
			high--
		}
		if low <= high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}
	return low
}
