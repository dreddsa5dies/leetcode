package main

func main() {}

/*
Runtime: 9 ms, faster than 6.64% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 72.79% of Go online submissions for Search in Rotated Sorted Array.
*/
// binary search
func search(nums []int, target int) int {
	var lo, hi int = 0, len(nums) - 1

	for lo <= hi {
		var mid = lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}

		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target <= nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
