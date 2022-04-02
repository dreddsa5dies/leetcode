/*
Runtime: 9 ms, faster than 55.62% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 4.3 MB, less than 98.28% of Go online submissions for Remove Duplicates from Sorted Array.
*/
package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) == 5)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
