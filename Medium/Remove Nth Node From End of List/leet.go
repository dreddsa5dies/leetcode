/*
Runtime: 3 ms, faster than 38.30% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 52.01% of Go online submissions for Remove Nth Node From End of List.
*/
package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{-1, head}
	cur, prevOfRemoval := dummyHead, dummyHead

	for cur.Next != nil {
		// n step delay for prevOfRemoval
		if n <= 0 {
			prevOfRemoval = prevOfRemoval.Next
		}
		cur = cur.Next
		n -= 1
	}

	// Remove the N-th node from end of list
	nthNode := prevOfRemoval.Next
	prevOfRemoval.Next = nthNode.Next

	return dummyHead.Next
}
