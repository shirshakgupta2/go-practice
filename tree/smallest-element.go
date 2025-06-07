package main


func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	currPosition := root

	for currPosition != nil || len(stack) > 0 {
		for currPosition != nil {
			stack = append(stack, currPosition)
			currPosition = currPosition.left
		}

		currPosition = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return currPosition.val
		}
		currPosition = currPosition.right
	}

	return -1 // If k is invalid
}
