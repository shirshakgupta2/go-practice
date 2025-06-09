package main

import "math"

func IsBalancedBTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.left)
	rightHeight := height(root.right)

	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		return false
	}

	return IsBalancedBTree(root.left) && IsBalancedBTree(root.right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.left), height(node.right))
}


func isBalanced(root *TreeNode) bool {
    ans:= true
    inorderUtil(root,&ans)
    return ans
}

func inorderUtil(root *TreeNode,ans *bool) {
    if root==nil{
        return 
    }
    inorderUtil(root.left, ans)
    if math.Abs(float64(HeightOfNode(root.left)-HeightOfNode(root.right))) > 1 {
        *ans = *ans && false
        return
    }
    inorderUtil(root.right, ans)

}


func HeightOfNode(root *TreeNode) int{
    if root ==nil{
        return 0
    }
    return max(HeightOfNode(root.left),HeightOfNode(root.right))+1

}