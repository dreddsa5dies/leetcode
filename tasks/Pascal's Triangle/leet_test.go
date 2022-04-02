// go test -v leet_test.go leet.go
package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	var tests = []struct {
		input int
		want  [][]int
	}{
		{1, [][]int{[]int{1}}},
		{5, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}},
	}
	for _, test := range tests {
		if got := generate(test.input); reflect.DeepEqual(got, test.want) {
			t.Errorf("generate(%v) = %v", test.input, got)
		}
	}
}
