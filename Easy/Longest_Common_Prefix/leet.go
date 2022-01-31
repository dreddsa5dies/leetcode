/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.3 MB, less than 99.47% of Go online submissions for Longest Common Prefix.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	if strs[0] == "" {
		return ""
	}

	s := ""
	count := 0
	for i := 0; i <= len(strs[0])-1; i++ {
		temp := 0
		for j := 1; j <= len(strs)-1; j++ {
			if i > len(strs[j]) {
				continue
			}
			if strings.HasPrefix(string(strs[j][i:]), string(strs[0][i])) {
				temp++
			}
			if temp == len(strs)-1 && count == i {
				s += string(strs[0][i])
				count++
			}
		}
	}

	return s
}
