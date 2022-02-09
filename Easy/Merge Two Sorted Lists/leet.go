/*
Runtime: 3 ms, faster than 40.75% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 96.50% of Go online submissions for Merge Two Sorted Lists.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var input1, input2, first, second ListNode
	input1 = ListNode{Val: 1, Next: &input2}
	input2 = ListNode{Val: 4}
	first = ListNode{Val: 3, Next: &second}
	second = ListNode{Val: 5}
	fmt.Println(mergeTwoLists(&input1, &first))
}

func moveNode(dest **ListNode, source **ListNode) {
	newNode := *source
	if newNode != nil {
		*source = newNode.Next
		newNode.Next = *dest
		*dest = newNode
	}
}

func mergeTwoLists(fSorted *ListNode, sSorted *ListNode) *ListNode {
	var result *ListNode = nil

	lastPtrRef := &result

	for {
		if fSorted == nil {
			*lastPtrRef = sSorted
			break
		} else if sSorted == nil {
			*lastPtrRef = fSorted
			break
		}
		if fSorted.Val <= sSorted.Val {
			moveNode(lastPtrRef, &fSorted)
		} else {
			moveNode(lastPtrRef, &sSorted)
		}
		lastPtrRef = &((*lastPtrRef).Next)
	}
	return result
}
