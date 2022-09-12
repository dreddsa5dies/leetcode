/*
Runtime: 236 ms, faster than 28.57% of Go online submissions for Contains Duplicate II.
Memory Usage: 13.7 MB, less than 37.71% of Go online submissions for Contains Duplicate II.
*/
package main

func main() {}

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]int)

	for i, v := range nums {
		if ii, ok := cache[v]; ok && (i-ii) <= k {
			return true
		}

		cache[v] = i
	}

	return false
}
