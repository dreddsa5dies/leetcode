// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_validPalindrome() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	// Output:
	// true
	// false
}
