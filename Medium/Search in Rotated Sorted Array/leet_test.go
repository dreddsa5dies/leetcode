// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_search() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	// Output:
	// 4
	// -1
}
