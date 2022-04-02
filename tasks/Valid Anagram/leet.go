/*
Runtime: 12 ms, faster than 38.33% of Go online submissions for Valid Anagram.
Memory Usage: 3 MB, less than 36.98% of Go online submissions for Valid Anagram.
*/
package main

import (
	"bytes"
	"sort"
)

func main() {}

// сделаем из строк срез байт, отсортируем и сравним
func isAnagram(s string, t string) bool {
	sb, tb := []byte(s), []byte(t)
	sort.Slice(sb, func(p, q int) bool { return sb[p] < sb[q] })
	sort.Slice(tb, func(p, q int) bool { return tb[p] < tb[q] })
	return bytes.Equal(sb, tb)
}

// в основном предлагаю решения через hashmap или перебор символов
