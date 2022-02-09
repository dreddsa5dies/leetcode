/*
Runtime: 1015 ms, faster than 5.12% of Go online submissions for Valid Palindrome.
Memory Usage: 54 MB, less than 5.12% of Go online submissions for Valid Palindrome.
*/
package main

import (
	"strings"
)

func main() {

}

func isPalindrome(s string) bool {
	fin := true

	s = normalize(s)

	if len(s) <= 1 {
		return true
	}

	if len(s)%2 == 0 {
		j := len(s) - 1
		for i := 0; i <= len(s)/2; i++ {
			if s[i] == s[j] {
				j--
				continue
			} else {
				return false
			}
		}
	} else {
		pivot := len(s)/2 + 1
		j := len(s) - 1
		for i := 0; i < pivot; i++ {
			if s[i] == s[j] {
				j--
				continue
			} else {
				return false
			}
		}
	}

	return fin
}

func normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	t := ""
	// 48--57, 97--122
	for _, v := range s {
		if (int(v) >= 48 && int(v) <= 57) || (int(v) >= 97 && int(v) <= 122) {
			t += string(v)
		}
	}

	return t
}
