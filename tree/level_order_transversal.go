package main

func LevelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var arr []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue)>0{
		temp :=queue[0]
		queue=queue[1:]
		arr = append(arr, temp.val)
		if temp.left!=nil{
			queue = append(queue, temp.left)
		}
		if temp.right!=nil{
			queue = append(queue, temp.right)
		}
	}
	return arr
}

func LevelOrderTraversalWithDepth(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var arr [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for {
		if len(queue)==0{
			break
		}
		size:=len(queue)
		var tempArr []int
		for size>0{
			temp := queue[0]
			queue=queue[1:]
			tempArr=append(tempArr, temp.val)
			if temp.left!=nil{
			queue = append(queue, temp.left)
			}
			if temp.right!=nil{
				queue = append(queue, temp.right)
			}
			size--
		}
		arr=append(arr,tempArr)
	}
	return arr
}

func ZigzagLevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var arr [][]int
	var queue []*TreeNode
	var isEvenLevel bool 
	queue = append(queue, root)
	isEvenLevel=false
	for {
		if len(queue)==0{
			break
		}
		var tempArr []int
		size:=len(queue)
		for size >0{
			temp:=queue[0]
			queue=queue[1:]
			
			if isEvenLevel{
				tempArr=append([]int{temp.val},tempArr...)
			}else{
				tempArr = append(tempArr, temp.val)
			}


			if temp.left!=nil{
				queue = append(queue, temp.left)
			}
			if temp.right !=nil{
				queue = append(queue, temp.right)
			}
			size--
		}
		arr = append(arr, tempArr)
		isEvenLevel=!isEvenLevel
	}

	return arr


	// if root == nil {
	// 	return [][]int{}
	// }

	// var arr [][]int
	// var queue []*TreeNode
	// queue = append(queue, root)
	// leftToRight := true

	// for len(queue) > 0 {
	// 	size := len(queue)
	// 	var tempArr []int

	// 	for size > 0 {
	// 		temp := queue[0]
	// 		queue = queue[1:]

	// 		if leftToRight {
	// 			tempArr = append(tempArr, temp.val)
	// 		} else {
	// 			tempArr = append([]int{temp.val}, tempArr...)
	// 		}

	// 		if temp.left != nil {
	// 			queue = append(queue, temp.left)
	// 		}
	// 		if temp.right != nil {
	// 			queue = append(queue, temp.right)
	// 		}
	// 		size--
	// 	}

	// 	arr = append(arr, tempArr)
	// 	leftToRight = !leftToRight
	// }

	// return arr
}