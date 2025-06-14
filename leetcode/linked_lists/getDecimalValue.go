package main


// func getDecimalValue(head *ListNode) int {
// 	result := 0
// 	for head != nil {
// 		result = (result << 1) | head.Val
// 		head = head.Next
// 	}
// 	return result
// }

func getDecimalValue(head *ListNode) int {
    if head==nil{
        return 0
    }
    var s string
    for head!=nil{
        if head.Val==1{
            s+=string('1')
        }else{
            s+=string('0')
        }
        head=head.Next
    }
    return binaryToDecimal(s)
}

func binaryToDecimal(s string) int{
    dec,base:=0,1
    for i:=len(s)-1;i>=0;i--{
        if string(s[i]) =="1"{
            dec += base 
        }
        base *=2
    }
    return dec

} 