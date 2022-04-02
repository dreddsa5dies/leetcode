/*
Runtime: 92 ms, faster than 96.19% of Go online submissions for Maximum Subarray.
Memory Usage: 8.5 MB, less than 99.68% of Go online submissions for Maximum Subarray.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(6 == maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(1 == maxSubArray([]int{1}))
	fmt.Println(23 == maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		max = int(math.Max(float64(nums[i]), float64(max)+float64(nums[i])))
		sum = int(math.Max(float64(sum), float64(max)))
	}
	return sum

}
