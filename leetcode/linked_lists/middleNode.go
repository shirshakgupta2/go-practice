package main



// func middleNode(head *ListNode) *ListNode {
// 	slow := head
// 	fast := head
// 	var prev *ListNode
// 	for fast != nil && fast.Next != nil {
// 		fast = fast.Next.Next
// 		next := slow.Next
// 		slow.Next = prev
// 		prev = slow
// 		slow = next
// 	}
// 	if fast == nil {
// 		// Even length
// 		return slow
// 	}
// 	// Odd length
// 	return slow
// }

func middleNode(head *ListNode) *ListNode {
   if head==nil{
        return head
    }
    var s []*ListNode
    for head!=nil{
        s = append(s, head)
        head=head.Next
    }
    return s[(len(s)/2)]
}

func sortArrayByParity(nums []int) []int {
    var even,odd []int
    for _,data:=range nums{
        if data%2==0{
            even= append(even,data)
        }else{
            odd=append(odd,data)
        }
    }
    return append(even,odd...)
}