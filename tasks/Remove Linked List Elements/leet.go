package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	for n := head; n != nil; n = n.Next {
		if n.Val == val {
			if prev == nil {
				head = n.Next
				continue
			}
			prev.Next = n.Next
			continue
		}
		prev = n
	}
	return head
}
