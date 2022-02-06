/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
Memory Usage: 1.8 MB, less than 98.42% of Go online submissions for Climbing Stairs.
*/
package main

func main() {
	println(climbStairs(1))
}

func climbStairs(n int) int {
	pre, cur, temp := 1, 1, 0

	for i := 1; i < n; i++ {
		temp = cur
		cur = cur + pre
		pre = temp
	}

	return cur
}
