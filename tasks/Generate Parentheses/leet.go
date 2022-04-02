/*
Runtime: 3 ms, faster than 55.23% of Go online submissions for Generate Parentheses.
Memory Usage: 2.8 MB, less than 82.01% of Go online submissions for Generate Parentheses.
*/
package main

func main() {}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	current := make([]byte, n*2)
	rec22(&ans, n, 0, 0, current)
	return ans
}

func rec22(ans *[]string, n int, left int, right int, current []byte) {
	if left+right == n*2 {
		*ans = append(*ans, string(current))
		return
	}

	if left < n {
		current[left+right] = '('
		rec22(ans, n, left+1, right, current)
	}

	if right < left {
		current[left+right] = ')'
		rec22(ans, n, left, right+1, current)
	}
}
