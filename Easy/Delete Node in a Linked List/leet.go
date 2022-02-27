/*
Runtime: 5 ms, faster than 42.09% of Go online submissions for Delete Node in a Linked List.
Memory Usage: 2.9 MB, less than 85.16% of Go online submissions for Delete Node in a Linked List.
*/
package main

func main() {}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// так как мы не имеем доступа к предыдущей ноде, то мы присваиваем значение следующей ноды этой и ставим ссылку на послеследующую, таким образом адрес ноды остается переменным, но фактически ее значение удаляется
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val

	next := node.Next
	node.Next = next.Next
	next.Next = nil
}
