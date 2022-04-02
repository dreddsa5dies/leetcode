package main

import (
	"strings"
)

func main() {}

/*
Runtime: 3 ms, faster than 34.49% of Go online submissions for Length of Last Word.
Memory Usage: 2.1 MB, less than 43.55% of Go online submissions for Length of Last Word.
*/
func lengthOfLastWord(s string) int {
	str := strings.Fields(s)
	return len(str[len(str)-1])
}
