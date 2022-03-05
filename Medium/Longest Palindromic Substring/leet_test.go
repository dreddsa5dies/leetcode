// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"babad", "bab"},
		{"cbbd", "cbbd"},
	}
	for _, test := range tests {
		if got := longestPalindrome(test.input); got != test.want {
			t.Errorf("longestPalindrome(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
