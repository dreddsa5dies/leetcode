// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{8, 2},
		{4, 2},
	}
	for _, test := range tests {
		if got := mySqrt(test.input); got != test.want {
			t.Errorf("mySqrt(%v) = %v", test.input, got)
		}
	}
}
