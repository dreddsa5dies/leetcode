package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 4 ms, faster than 76.25% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.2 MB, less than 42.78% of Go online submissions for Remove Duplicates from Sorted List.
*/
// recursion
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}

/*
Runtime: 8 ms, faster than 25.75% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.1 MB, less than 42.78% of Go online submissions for Remove Duplicates from Sorted List.
*/
// // brute
// func deleteDuplicates(head *ListNode) *ListNode {

//     cur := head
//     for cur != nil {
//         if cur.Next != nil && cur.Val == cur.Next.Val {
//             cur.Next = cur.Next.Next
//             continue
//         }
//         cur = cur.Next
//     }

//     return head
// }
