package main

func main() {}

/*
Runtime: 4 ms, faster than 60.72% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 15.07% of Go online submissions for Search Insert Position.
*/
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
