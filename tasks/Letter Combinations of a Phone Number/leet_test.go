// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_letterCombinations() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	// Output:
	// [ad ae af bd be bf cd ce cf]
	// []
	// [a b c]
}
