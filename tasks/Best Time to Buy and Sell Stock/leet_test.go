// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
	}
	for _, test := range tests {
		if got := maxProfit(test.input); got != test.want {
			t.Errorf("maxProfit(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
