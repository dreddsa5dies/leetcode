/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
Memory Usage: 2.3 MB, less than 70.67% of Go online submissions for Divide Two Integers.
*/
package main

func main() {}

func divide(dividend int, divisor int) int {
	if dividend == (-2<<30) && divisor == -1 {
		return (2 << 30) - 1
	}

	dividendAbs := abs(dividend)
	divisorAbs := abs(divisor)
	res := 0

	for dividendAbs-divisorAbs >= 0 {
		exponent := 0
		for dividendAbs-divisorAbs<<exponent >= 0 {
			exponent++
		}
		res += 1 << (exponent - 1)
		dividendAbs -= divisorAbs << (exponent - 1)
	}
	if (divisor > 0) == (dividend >= 0) {
		return res
	}
	return -res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
