/*
Runtime: 3 ms, faster than 33.76% of Go online submissions for Merge Sorted Array.
Memory Usage: 2.3 MB, less than 76.96% of Go online submissions for Merge Sorted Array.
*/
package main

import "sort"

func main() {
	a := make([]int, 6)
	a[0], a[1], a[3] = 1, 2, 3
	merge(a, 3, []int{2, 3, 4}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}
	sort.Ints(nums1)
}
