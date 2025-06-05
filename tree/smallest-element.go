package main


func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.val
		}
		curr = curr.right
	}

	return -1 // If k is invalid
}
