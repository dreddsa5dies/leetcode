/*
Runtime: 37 ms, faster than 42.39% of Go online submissions for Find Pivot Index.
Memory Usage: 6.3 MB, less than 69.54% of Go online submissions for Find Pivot Index.
*/
package main

func main() {}

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	s := 0
	for i, num := range nums {
		if sum-num-s == s {
			return i
		}
		s += num
	}

	return -1
}
