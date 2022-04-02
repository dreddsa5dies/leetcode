package main

func main() {}

/*
Runtime: 7 ms, faster than 69.86% of Go online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 4 MB, less than 21.67% of Go online submissions for Find First and Last Position of Element in Sorted Array.
*/
// бинарный поиск
// Time complexity: O(log n) where n is the length of nums
// Space complexity: O(1)
func searchRange(nums []int, target int) []int {
	bisectLeft := func(arr []int, target int) int {
		left, right := 0, len(arr)

		for left < right {
			mid := left + (right-left)/2

			if target > arr[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	bisectRight := func(arr []int, target int) int {
		left, right := 0, len(arr)

		for left < right {
			mid := left + (right-left)/2

			if target < arr[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := bisectLeft(nums, target)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := bisectRight(nums, target)
	return []int{left, right - 1}
}

// /*
// Runtime: 9 ms, faster than 52.59% of Go online submissions for Find First and Last Position of Element in Sorted Array.
// Memory Usage: 4.2 MB, less than 8.14% of Go online submissions for Find First and Last Position of Element in Sorted Array.
// */
// // простой перебор
// // Time complexity: O(n) where n is the length of nums
// func searchRange(nums []int, target int) []int {
// 	tmp := []int{}
// 	for k, v := range nums {
// 		if v == target {
// 			tmp = append(tmp, k)
// 		}
// 	}
// 	if len(tmp) == 0 {
// 		return []int{-1, -1}
// 	} else {
// 		return []int{tmp[0], tmp[len(tmp)-1]}
// 	}
// }
