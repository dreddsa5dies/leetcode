package main

/*
Runtime: 17 ms, faster than 23.16% of Go online submissions for Linked List Cycle.
Memory Usage: 6.6 MB, less than 6.21% of Go online submissions for Linked List Cycle.
*/
func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	iter := head
	visited := map[*ListNode]bool{}
	for iter != nil {
		if visited[iter] {
			return true
		}
		visited[iter] = true
		iter = iter.Next
	}
	return false
}
