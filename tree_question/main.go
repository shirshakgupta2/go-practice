package main

import "fmt"

type TreeNode struct {
	Data      int
	RightNode *TreeNode
	LeftNode  *TreeNode
}

func buildBinaryTree(treeValues []int) *TreeNode {
	var root *TreeNode
	for _, data := range treeValues {
		root = insertDataToBT(root, data)
	}

	return root
}

func insertDataToBT(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Data: value}
	}
	if value > root.Data {
		root.RightNode = insertDataToBT(root.RightNode, value)
	} else {
		root.LeftNode = insertDataToBT(root.LeftNode, value)
	}
	return root
}

func NewTree() *TreeNode{
	return &TreeNode{}
}


func main() {

	values := []int{10, 5, 15, 3, 7, 12, 18}
	root := buildBinaryTree(values)
	fmt.Println("ROOT:",root)


	//BASIC INSERTION OF DATA IN NODE
	treeRoot:=NewTree()
	treeRoot.BasicInsertion()
	PrintInOrderTransversal(treeRoot)
	fmt.Println()
	PrintPreOrderTransversal(treeRoot)
	fmt.Println()
	PrintPostOrderTransversal(treeRoot)
}
