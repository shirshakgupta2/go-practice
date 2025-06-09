package main

import "fmt"

// Inorder traversal (for validation/debugging)
func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Print(root.val, " ")
	InOrderTraversal(root.right)
}
