/*
Runtime: 60 ms, faster than 5.42% of Go online submissions for Palindrome Number.
Memory Usage: 6.2 MB, less than 18.56% of Go online submissions for Palindrome Number.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1211))
}

func isPalindrome(x int) bool {
	fin := true
	switch {
	case x < 0:
		return false
	case x >= 0 && x < 10:
		return true
	}
	tmp := []int{}
	counter := x
	a := 10
	for i := 1; i <= counter; i *= 10 {
		tmp = append(tmp, x%a)
		x = (x - x%a) / 10
	}

	var pivot int
	if len(tmp)%2 == 0 {
		pivot = (len(tmp) - 1) / 2
		tmp1 := make([]int, pivot+1)
		tmp2 := make([]int, pivot+1)
		tmp1, tmp2 = tmp[0:pivot+1], tmp[pivot+1:]
		for i, j := 0, len(tmp2)-1; i < j; i, j = i+1, j-1 {
			tmp2[i], tmp2[j] = tmp2[j], tmp2[i]
		}
		for k, _ := range tmp1 {
			if tmp1[k] != tmp2[k] {
				fin = false
				break
			}
		}
	} else {
		pivot = (len(tmp)) / 2
		tmp1 := make([]int, pivot+1)
		tmp2 := make([]int, pivot+1)
		tmp1, tmp2 = tmp[0:pivot], tmp[pivot+1:]
		for i, j := 0, len(tmp2)-1; i < j; i, j = i+1, j-1 {
			tmp2[i], tmp2[j] = tmp2[j], tmp2[i]
		}
		for k, _ := range tmp1 {
			if tmp1[k] != tmp2[k] {
				fin = false
				break
			}
		}
	}
	return fin
}
