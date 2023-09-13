package main

import "testing"

func TestTwoSum(t *testing.T) {
	t.Run("Example Test Cases", func(t *testing.T) {
		t.Run("(1)", func(t *testing.T) {
			nums := []int{2, 7, 11, 15}
			target := 9
			result := []int{0, 1}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("expected result: %v != %v - real result", output, result)
			}
		})
		t.Run("(2)", func(t *testing.T) {
			nums := []int{3, 2, 4}
			target := 6
			result := []int{1, 2}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("expected result: %v != %v - real result", output, result)
			}
		})
		t.Run("(3)", func(t *testing.T) {
			nums := []int{3, 3}
			target := 6
			result := []int{0, 1}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("expected result: %v != %v - real result", output, result)
			}
		})
	})
	t.Run("Some Additional Test Cases", func(t *testing.T) {
		t.Run("(1)", func(t *testing.T) {
			nums := []int{3, 1, 4, 0, 5, 9, 2, 18}
			target := 10
			result := []int{1, 5}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("%v != %v", result, output)
			}
		})
		t.Run("(2)", func(t *testing.T) {
			nums := []int{-1, 0, 7, 2, 3, 4, 26, -80, -2}
			target := 0
			result := []int{3, 8}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("%v != %v", result, output)
			}
		})
		t.Run("(3)", func(t *testing.T) {
			nums := []int{50, 100, -5000, 1500, 156165, 156165, 18456, 152397, 1561, 8985, -5000, -10000}
			target := -10000
			result := []int{2, 10}
			output := twoSum(nums, target)

			if !isSlicesEquals2(result, output) {
				t.Errorf("%v != %v", result, output)
			}
		})
	})
}
