// go test -v leet_test.go leet.go
package main

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}
	for _, test := range tests {
		if got := titleToNumber(test.input); got != test.want {
			t.Errorf("titleToNumber(%v) = %v", test.input, got)
		}
	}
}
