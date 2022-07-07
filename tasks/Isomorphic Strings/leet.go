/*
Runtime: 12 ms, faster than 21.10% of Go online submissions for Isomorphic Strings.
Memory Usage: 3.9 MB, less than 12.77% of Go online submissions for Isomorphic Strings.
*/
package main

func main() {}

func isIsomorphic(s string, t string) bool {
	sPattern, tPattern := make(map[uint8]int, len(s)), make(map[uint8]int, len(t))
	for index := range s {
		if sPattern[s[index]] != tPattern[t[index]] {
			return false
		} else {
			sPattern[s[index]] = index + 1
			tPattern[t[index]] = index + 1
		}
	}

	return true
}
