// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{321, 123},
		{120, 21},
		{-123, -321},
		{1534236469, 0},
	}
	for _, test := range tests {
		if got := reverse(test.input); got != test.want {
			t.Errorf("reverse(%v) = %v, need %v", test.input, got, test.want)
		}
	}
}
