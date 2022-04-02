/*
Runtime: 2 ms, faster than 50.00% of Go online submissions for Excel Sheet Column Number.
Memory Usage: 2.1 MB, less than 78.13% of Go online submissions for Excel Sheet Column Number.
*/
package main

import "strings"

func main() {}

func titleToNumber(columnTitle string) int {
	var foo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	count := 0
	for _, v := range columnTitle {
		count *= 26
		count += strings.Index(foo, string(v)) + 1
	}
	return count
}
