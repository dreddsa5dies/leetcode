// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{19, true},
		{2, false},
	}
	for _, test := range tests {
		if got := isHappy(test.input); got != test.want {
			t.Errorf("isHappy(%v) = %v", test.input, got)
		}
	}
}
