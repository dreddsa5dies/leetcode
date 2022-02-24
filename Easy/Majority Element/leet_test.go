// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, test := range tests {
		if got := majorityElement(test.input); got != test.want {
			t.Errorf("majorityElement(%v) = %v", test.input, got)
		}
	}
}
