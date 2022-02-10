/*
Runtime: 31 ms, faster than 32.50% of Go online submissions for Single Number.
Memory Usage: 6.7 MB, less than 32.03% of Go online submissions for Single Number.
*/
package main

func main() {

}

func singleNumber(nums []int) int {
	fin := map[int]int{}
	var exp int

	for k := range nums {
		fin[nums[k]]++
	}

	for key, val := range fin {
		if val == 1 {
			exp = key
		}
	}

	return exp
}
