package main

func main() {}

// definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 3 ms, faster than 29.78% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2 MB, less than 42.08% of Go online submissions for Binary Tree Preorder Traversal.
*/
func preorderTraversal(root *TreeNode) []int {
	fin := []int{}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node != nil {

			fin = append(fin, node.Val)

			preorder(node.Left)
			preorder(node.Right)
		}
	}

	preorder(root)

	return fin
}
