/*
Runtime: 36 ms, faster than 75.08% of Go online submissions for Intersection of Two Linked Lists.
Memory Usage: 7.7 MB, less than 31.15% of Go online submissions for Intersection of Two Linked Lists.
*/
package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := map[*ListNode]bool{}
	for currA := headA; currA != nil; currA = currA.Next {
		hashMap[currA] = true
	}

	for currB := headB; currB != nil; currB = currB.Next {
		_, k := hashMap[currB]
		if k {
			return currB
		}
	}

	return nil
}
