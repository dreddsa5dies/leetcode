/*
Runtime: 7 ms, faster than 16.16% of Go online submissions for Running Sum of 1d Array.
Memory Usage: 2.7 MB, less than 89.74% of Go online submissions for Running Sum of 1d Array.
*/
package main

func main() {}

func runningSum(nums []int) []int {
	for k := range nums {
		if k == len(nums)-1 {
			break
		}
		nums[k+1] = nums[k+1] + nums[k]
	}
	return nums
}
