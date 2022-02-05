/*
Runtime: 11 ms, faster than 26.18% of Go online submissions for Sqrt(x).
Memory Usage: 2.2 MB, less than 86.68% of Go online submissions for Sqrt(x).
*/
package main

import "fmt"

func main() {
	fmt.Println(mySqrt(2))
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	} else {
		smallCandidate := mySqrt(x>>2) << 1
		largeCandidate := smallCandidate + 1
		if largeCandidate*largeCandidate > x {
			return smallCandidate
		} else {
			return largeCandidate
		}
	}
}
