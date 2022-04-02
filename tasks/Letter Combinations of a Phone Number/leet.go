package main

func main() {}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
Memory Usage: 2 MB, less than 94.89% of Go online submissions for Letter Combinations of a Phone Number.
*/
// recursion rule
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	data := map[uint8][]string{
		uint8(50): []string{"a", "b", "c"},
		uint8(51): []string{"d", "e", "f"},
		uint8(52): []string{"g", "h", "i"},
		uint8(53): []string{"j", "k", "l"},
		uint8(54): []string{"m", "n", "o"},
		uint8(55): []string{"p", "q", "r", "s"},
		uint8(56): []string{"t", "u", "v"},
		uint8(57): []string{"w", "x", "y", "z"},
	}

	if len(digits) == 1 {
		return data[digits[0]]
	}

	res := []string{}

	backtrack("", &res, 0, digits, data)
	return res
}

func backtrack(curStr string, res *[]string, index int, digits string, digitsMap map[uint8][]string) {
	if index == len(digits) {
		*res = append(*res, curStr)
		return
	}
	for _, ch := range digitsMap[digits[index]] {
		backtrack(curStr+ch, res, index+1, digits, digitsMap)
	}
}
