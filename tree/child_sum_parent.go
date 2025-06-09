package main



func ChildSumParent(root *TreeNode) bool {
    if root==nil{
        return true
    }
    if root.left==nil && root.right==nil{
        return true
    }
    var l,r int
    if root.left!=nil{
        l=root.left.val
    }
    if root.right!=nil{
        r=root.right.val
    }
    if  root.val==l+r && ChildSumParent(root.right) && ChildSumParent(root.left) {
        return true
    }
    return false

}