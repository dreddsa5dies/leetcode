/*
Runtime: 12 ms, faster than 71.76% of Go online submissions for Add Two Numbers.
Memory Usage: 4.7 MB, less than 67.41% of Go online submissions for Add Two Numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// думал, что надо получить число, сложить, развернуть в обратном порядке и выдать как связанный список, а тут проще!!!!!!!

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, sum := new(ListNode), 0
	for cur := dummy; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummy.Next
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := reverse(getNum(l1) + getNum(l2))
	var fin = new(ListNode)

	x := tmp
	a := 10
	for i := 1; i <= tmp; i *= 10 {
		if i == 1 {
			fin = &ListNode{x % a, nil}
			x = (x - x%a) / 10
			continue
		}

		t := &ListNode{x % a, nil}
		t.Next = fin
		fin = t

		x = (x - x%a) / 10
	}

	return fin
}

func reverse(value int) int {

	intString := strconv.Itoa(value)

	newString := ""

	for x := len(intString); x > 0; x-- {
		newString += string(intString[x-1])
	}

	newInt, err := strconv.Atoi(newString)

	if err != nil {
		fmt.Println("Error converting string to int")
	}

	return newInt
}

func getNum(l *ListNode) int {
	tmp := []int{}

	for l != nil {
		tmp = append(tmp, l.Val)
		l = l.Next
	}

	fin, d := 0, 1

	for i := len(tmp) - 1; i >= 0; i-- {
		fin += tmp[i] * d
		d *= 10
	}

	return fin
}
