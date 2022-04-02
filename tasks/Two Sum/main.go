/*
Runtime: 40 ms, faster than 20.15% of Go online submissions for Two Sum.
Memory Usage: 3.6 MB, less than 99.49% of Go online submissions for Two Sum.
*/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2}, 2))
}

func twoSum(nums []int, target int) []int {
	tmp := make([]int, 2)
	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				tmp[0], tmp[1] = i, j
			}
		}
	}
	return tmp
}
