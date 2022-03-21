// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_searchRange() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	// Output:
	// [3 4]
	// [-1 -1]
}
