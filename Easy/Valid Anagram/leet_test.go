// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		want   bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, test := range tests {
		if got := isAnagram(test.input1, test.input2); got != test.want {
			t.Errorf("isAnagram(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
