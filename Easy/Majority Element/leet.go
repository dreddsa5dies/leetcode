/*
Runtime: 25 ms, faster than 42.43% of Go online submissions for Majority Element.
Memory Usage: 6.2 MB, less than 44.26% of Go online submissions for Majority Element.
*/
package main

func main() {}

func majorityElement(nums []int) int {
	visited := map[int]int{}

	for k := range nums {
		visited[nums[k]]++
	}

	var max, maj int
	for k, v := range visited {
		if max < v {
			max = v
			maj = k
		}
	}

	return maj
}
