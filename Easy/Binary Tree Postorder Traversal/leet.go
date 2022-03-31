package main

func main() {}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2 MB, less than 39.47% of Go online submissions for Binary Tree Postorder Traversal.
*/
func postorderTraversal(root *TreeNode) []int {
	fin := []int{}

	var postorder func(node *TreeNode)

	postorder = func(node *TreeNode) {
		if node != nil {
			postorder(node.Left)
			postorder(node.Right)

			fin = append(fin, node.Val)
		}
	}

	postorder(root)

	return fin
}
