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



func main() {
	values := []int{10, 5, 15, 3, 7, 12, 18,9,8,19}
	root := buildBST(values)

	//! In-order traversal to print the BST
	fmt.Println("In-order Traversal: ")
	InOrderTraversal(root)
	fmt.Println()

	//! Level order traversal to print the BST
	fmt.Println("Level Order Traversal: ", LevelOrderTraversal(root))

	//! Level order traversal with depth
	fmt.Println("Level Order Traversal with Depth: ", LevelOrderTraversalWithDepth(root))

	//! Zigzag level order traversal
	fmt.Println("Zigzag Level Order Traversal: ", ZigzagLevelOrderTraversal(root))

	k := 2
	// k := 0
	// fmt.Print("Enter k: ")
	//! Uncomment the following lines to take user input for k
	// _, err := fmt.Scan(&k)
	// if err != nil || k <= 0 {
	// 	fmt.Println("k should be a positive integer")
	// 	return
	// }
	fmt.Println("Kth smallest element:", kthSmallest(root, k))
	//! child sum parent
	fmt.Println("Child Sum Parent:", ChildSumParent(root))

	//! Check if the tree is balanced
	fmt.Println("Is Balanced:", IsBalancedBTree(root))
	fmt.Println("Is Balanced (Inorder):", isBalanced(root))

	//!Right view of the tree
	fmt.Println("Right View of the Tree:", rightViewOfTree(root))

	//!Left view of the tree
	fmt.Println("Left View of the Tree:", leftViewOfTree(root))


}
