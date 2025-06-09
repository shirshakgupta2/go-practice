package main

import "fmt"

func (root *TreeNode) BasicInsertion() {
	root.Data = 1
	root.LeftNode = &TreeNode{Data: 2}
	root.RightNode = &TreeNode{Data: 3}

	root.LeftNode.LeftNode = &TreeNode{Data: 4}
	root.LeftNode.RightNode = &TreeNode{Data: 5}
}


func PrintInOrderTransversal(root *TreeNode) {
// PrintInOrderTransversal performs an in-order traversal of the binary tree
	// and prints the data of each node in the order: left, root, right.
	if root == nil {
		return
	}
	// Traverse the left subtree
	PrintInOrderTransversal(root.LeftNode)
	// Visit the root node
	// Print the data of the current node
	fmt.Print(root.Data, " ")
	// Traverse the right subtree
	PrintInOrderTransversal(root.RightNode)
}


func PrintPreOrderTransversal(root *TreeNode) {
	// PrintPreOrderTransversal performs a pre-order traversal of the binary tree
	// and prints the data of each node in the order: root, left, right.
	if root == nil {
		return
	}
	// Visit the root node
	fmt.Print(root.Data, " ")
	// Traverse the left subtree
	PrintPreOrderTransversal(root.LeftNode)
	// Traverse the right subtree
	PrintPreOrderTransversal(root.RightNode)
}

func PrintPostOrderTransversal(root *TreeNode) {
	// PrintPostOrderTransversal performs a post-order traversal of the binary tree
	// and prints the data of each node in the order: left, right, root.
	if root == nil {
		return
	}
	// Traverse the left subtree
	PrintPostOrderTransversal(root.LeftNode)
	// Traverse the right subtree
	PrintPostOrderTransversal(root.RightNode)
	// Visit the root node
	fmt.Print(root.Data, " ")
}