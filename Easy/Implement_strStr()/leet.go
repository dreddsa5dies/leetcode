/*
Runtime: 1 ms, faster than 88.07% of Go online submissions for Implement strStr().
Memory Usage: 2.4 MB, less than 97.73% of Go online submissions for Implement strStr().
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(2 == strStr("hello", "ll"))
	fmt.Println(-1 == strStr("aaaaa", "bba"))
	fmt.Println(0 == strStr("aaaaa", "a"))
}

func strStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}

	return strings.Index(haystack, needle)
}
