package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
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
// c) порядок не важен:
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

// d) если порядок важен:
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
