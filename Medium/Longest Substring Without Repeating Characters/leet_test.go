// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.input); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
