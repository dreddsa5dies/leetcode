// go test -v leet_test.go leet.go
package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		input int
		want  []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, test := range tests {
		if got := fizzBuzz(test.input); reflect.DeepEqual(got, test.want) {
			t.Errorf("fizzBuzz(%v) = %v", test.input, got)
		}
	}
}
