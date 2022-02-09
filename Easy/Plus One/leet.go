/*
Runtime: 3 ms, faster than 23.97% of Go online submissions for Plus One.
Memory Usage: 2 MB, less than 94.20% of Go online submissions for Plus One.
*/
package main

import "fmt"

func main() {
	fmt.Println(equal([]int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1})), plusOne([]int{4, 3, 2, 1}))
	fmt.Println(equal([]int{1, 0}, plusOne([]int{9})), plusOne([]int{9}))
	fmt.Println(equal([]int{1, 0, 0}, plusOne([]int{9, 9})), plusOne([]int{9, 9}))
	fmt.Println(equal([]int{9, 1, 0}, plusOne([]int{9, 0, 9})), plusOne([]int{9, 0, 9}))
	fmt.Println(equal([]int{1, 9, 1, 0}, plusOne([]int{1, 9, 0, 9})), plusOne([]int{1, 9, 0, 9}))
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func plusOne(digits []int) []int {
	count := len(digits) - 1

	if valid(digits[count]) {
		digits[count]++
		return digits
	}

	for j := len(digits) - 1; j >= 0; j-- {
		if !valid(digits[j]) && count == j && j != 0 {
			digits[j] = 0
			count--
		} else if valid(digits[j]) && count == j {
			digits[j]++
			break
		} else if !valid(digits[j]) && j == 0 {
			digits[j] = 1
			digits = append(digits, 0)
		}
	}

	return digits
}

func valid(v int) bool {
	if v < 9 {
		return true
	}
	return false
}
