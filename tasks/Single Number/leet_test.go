// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
	for _, test := range tests {
		if got := singleNumber(test.input); got != test.want {
			t.Errorf("singleNumber(%v) = %v", test.input, got)
		}
	}
}
