package main

func main() {}

// рекурсивный метод
/*
Runtime: 4 ms, faster than 95.62% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.7 MB, less than 86.56% of Go online submissions for Longest Palindromic Substring.
*/
func longestPalindrome(s string) string {
	max, res := 0, ""

	findPalindrome := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > max {
				max = right - left + 1
				res = s[left : right+1]
			}
			left--
			right++
		}
	}

	for i := range s {
		findPalindrome(i, i)   // odd length
		findPalindrome(i, i+1) // even length
	}

	return res
}

/*
/*
Runtime: 4 ms, faster than 95.62% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.7 MB, less than 86.56% of Go online submissions for Longest Palindromic Substring.
/*
// метод перебора
func longestPalindrome(s string) string {
    if len(s) == 0 || len(s) == 1 { return s }
    lStart, maxLen := 0, 1
    for i := 0; i < len(s); {
        if len(s) - i <= maxLen/2 { break }
        // skip duplicates
        k, j := i, i
        for k < len(s)-1 && s[k]==s[k+1] { k++ }
        // expand around centre
        i = k+1
        for k < len(s)-1 && j > 0 && s[k+1] == s[j-1] { k++; j--;}
        newLen := k-j+1
        if newLen > maxLen { maxLen = newLen; lStart = j }
    }
    return s[lStart:lStart+maxLen]
}
*/
