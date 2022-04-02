/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
Memory Usage: 1.9 MB, less than 73.06% of Go online submissions for Number of 1 Bits.
*/
package main

import "math/bits"

func main() {}

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
