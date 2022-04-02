package main

import "math/big"

func main() {}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
Memory Usage: 2.2 MB, less than 84.90% of Go online submissions for Add Binary.
*/
func addBinary(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	return sum.Text(2)
}
