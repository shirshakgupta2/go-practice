package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Insert a new value into the BST
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

// Helper function to build BST from slice
func buildBST(values []int) *TreeNode {
	var root *TreeNode
	for _, v := range values {
		root = insert(root, v)
	}
	return root
}

// Inorder traversal (for validation/debugging)
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.val, " ")
	inorder(root.right)
}

func main() {
	values := []int{10, 5, 15, 3, 7, 12, 18}
	root := buildBST(values)

	fmt.Println("Inorder Traversal: ")
	inorder(root)
	fmt.Println()


	k := 2
	// k := 0
	// fmt.Print("Enter k: ")
	// _, err := fmt.Scan(&k)
	// if err != nil || k <= 0 {
	// 	fmt.Println("k should be a positive integer")
	// 	return
	// }
	fmt.Println("Kth smallest element:", kthSmallest(root, k))
}
