// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, test := range tests {
		if got := containsDuplicate(test.input); got != test.want {
			t.Errorf("containsDuplicate(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
