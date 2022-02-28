/*
Runtime: 25 ms, faster than 48.11% of Go online submissions for Missing Number.
Memory Usage: 6.6 MB, less than 25.05% of Go online submissions for Missing Number.
*/
package main

func main() {}

// поиск через общую сумму!!!
func missingNumber(nums []int) int {
	size := len(nums)
	var current, total int

	// Time: O(n)
	for _, num := range nums {
		if num <= size {
			current += num
		}
	}

	// Time: O(n)
	for i := 0; i < size+1; i++ {
		total += i
	}

	return total - current
}
