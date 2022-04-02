package main

func main() {}

/*
Runtime: 13 ms, faster than 50.50% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3.1 MB, less than 57.55% of Go online submissions for Longest Substring Without Repeating Characters.
*/
// решение через мапу
func lengthOfLongestSubstring(s string) int {
	lastOccurrenceMap := make(map[rune]int)
	startPosition, maxLength, currentLength := 0, 0, 0
	for currentPosition, char := range s {
		lastOccurrence, hasOccurred := lastOccurrenceMap[char]
		if hasOccurred && lastOccurrence >= startPosition {
			startPosition = lastOccurrence + 1
		}
		currentLength = currentPosition - startPosition + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
		lastOccurrenceMap[char] = currentPosition
	}
	return maxLength
}

/*
Runtime: 12 ms, faster than 57.41% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 2.7 MB, less than 82.01% of Go online submissions for Longest Substring Without Repeating Characters.
*/
/*
// https://www.code-recipe.com/post/longest-substring-without-repeating-characters
func lengthOfLongestSubstring(s string) int {
	// Length of the given input string
	n := len(s)

	// Length of longest substring
	var result int

	// Map to store visited characters along with their index
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			result = max(result, end-start)

			// Remove all characters before the new
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	result = max(result, n-start)

	return result
}

// max returns the max of num1 and num2
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
*/
