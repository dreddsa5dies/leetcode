/*
Runtime: 203 ms, faster than 61.88% of Go online submissions for Palindrome Linked List.
Memory Usage: 12 MB, less than 18.50% of Go online submissions for Palindrome Linked List.
*/
package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// стандартный мой подход. Остальные представленные в комментариях не такие элегантные и многословные
func isPalindrome(head *ListNode) bool {
	tmp := []int{}

	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}

	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}
	return true
}
