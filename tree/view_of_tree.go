package main


func rightViewOfTree(root *TreeNode) []int {
    if root ==nil{
        return []int{}
    }
    var queue []*TreeNode
    var arr []int
    queue=append(queue,root)
    for {
        if len(queue)==0{
            break
        }
        var tempArr []int
        size:=len(queue)
        for size>0{
            temp:=queue[0]
            tempArr=append(tempArr,temp.val)
            queue=queue[1:]
            if temp.left!=nil{
				queue=append(queue,temp.left)
            }
			if temp.right!=nil{
				queue=append(queue,temp.right)
			}
            size--
        }
        arr=append(arr,tempArr[len(tempArr)-1])
    }
    return arr
}
func leftViewOfTree(root *TreeNode) []int {
    if root ==nil{
        return []int{}
    }
    var queue []*TreeNode
    var arr []int
    queue=append(queue,root)
    for {
        if len(queue)==0{
            break
        }
        var tempArr []int
        size:=len(queue)
        for size>0{
            temp:=queue[0]
            tempArr=append(tempArr,temp.val)
            queue=queue[1:]
            if temp.left!=nil{
				queue=append(queue,temp.left)
            }
			if temp.right!=nil{
				queue=append(queue,temp.right)
			}
            size--
        }
        arr=append(arr,tempArr[0])
    }
    return arr
}