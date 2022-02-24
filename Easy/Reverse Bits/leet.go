/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Bits.
Memory Usage: 2.6 MB, less than 84.25% of Go online submissions for Reverse Bits.
*/
package main

import "math/bits"

func main() {}

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}
