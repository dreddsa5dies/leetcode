// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, test := range tests {
		if got := firstUniqChar(test.input); got != test.want {
			t.Errorf("firstUniqChar(%v) = %v", test.input, got)
		}
	}
}
