package main

func main() {}
/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Excel Sheet Column Title.
Memory Usage: 1.9 MB, less than 11.76% of Go online submissions for Excel Sheet Column Title.
*/
func convertToTitle(n int) string {
	chars := []byte{}

	for n > 0 {
		n--
		chars = append(chars, byte('A'+n%26))
		n /= 26
	}

	// reverse order
	for h, t := 0, len(chars)-1; h < t; h, t = h+1, t-1 {
		chars[h], chars[t] = chars[t], chars[h]
	}

	return string(chars)
}
