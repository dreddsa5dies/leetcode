package main

import "sort"

func main() {}

/*
Runtime: 28 ms, faster than 95.95% of Go online submissions for 3Sum.
Memory Usage: 7.5 MB, less than 77.58% of Go online submissions for 3Sum.
*/
// плавающее окно - O(n^2)
func threeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	sort.Ints(nums)

	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}

// - memory out
// 0(n^3) - brute
/*
func threeSum(nums []int) [][]int {
	n := len(nums)

	// Hashmap to store the triplets
	// Helps avoid duplicate triplets
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	// Iterate through the hashmap and add the triplets to the result array
	var result [][]int
	for _, triplet := range tripletMap {
		result = append(result, triplet)
	}
	return result
}
*/
