// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}
	for _, test := range tests {
		if got := maxArea(test.input); got != test.want {
			t.Errorf("maxArea(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
