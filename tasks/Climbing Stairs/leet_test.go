// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 2},
		{3, 3},
	}
	for _, test := range tests {
		if got := climbStairs(test.input); got != test.want {
			t.Errorf("climbStairs(%v) = %v", test.input, got)
		}
	}
}
