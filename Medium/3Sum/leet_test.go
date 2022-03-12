// go test -v leet_test.go leet.go
package main

import (
	"fmt"
)

func Example_threeSum() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	// Output:
	// [[-1 -1 2] [-1 0 1]]
	// []
	// []
}
