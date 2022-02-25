/*
Runtime: 3 ms, faster than 32.83% of Go online submissions for Happy Number.
Memory Usage: 2.1 MB, less than 72.64% of Go online submissions for Happy Number.
*/
package main

func main() {}

func isHappy(n int) bool {
	nMap := map[int]bool{}

	for {
		if n = getSquareSum(n); n == 1 {
			return true
		} else {
			if _, ok := nMap[n]; ok { // break if already in num Map
				return false
			} else {
				nMap[n] = true
			}
		}
	}
}

func getSquareSum(num int) (result int) {
	for num >= 1 {
		rem := num % 10
		result += rem * rem
		num /= 10
	}
	return
}
