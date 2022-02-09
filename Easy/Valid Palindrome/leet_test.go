// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{" ", true},
		{"race a car", false},
		{"A man, a plan, a canal: Panama", true},
		{"a.", true},
		{",.", true},
		{"ab@a", true},
		{"0P", false},
	}
	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%v) = %v", test.input, got)
		}
	}
}
