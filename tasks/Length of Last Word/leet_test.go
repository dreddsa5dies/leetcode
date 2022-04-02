// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_lengthOfLastWord() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	// Output:
	// 5
	// 4
	// 6
}
