/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
Memory Usage: 2.2 MB, less than 40.33% of Go online submissions for Remove Element.
*/
package main

import "fmt"

func main() {
	fmt.Println(2 == removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(5 == removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	tmp := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			tmp = append(tmp, nums[i])
		}
	}

	copy(nums, tmp)
	return len(tmp)
}
