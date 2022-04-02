package main

func main() {}

/*
Runtime: 34 ms, faster than 19.29% of Go online submissions for Valid Palindrome II.
Memory Usage: 10.4 MB, less than 5.14% of Go online submissions for Valid Palindrome II.
*/
func validPalindrome(s string) bool {
	return validPalindromeWithDelete(s, true)
}

func validPalindromeWithDelete(s string, canDelete bool) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return validPalindromeWithDelete(s[1:len(s)-1], canDelete)
	}
	if canDelete {
		return validPalindromeWithDelete(s[:len(s)-1], false) || validPalindromeWithDelete(s[1:], false)
	}

	return false
}
