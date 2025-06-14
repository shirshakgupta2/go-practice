package main

import "fmt"
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Example usage:
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	fmt.Println(getDecimalValue(head)) // Output: 5

	// Example usage:
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	mid := middleNode(head)
	fmt.Println(mid.Val) // Output: 3

	// Example usage: Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
	nums := []int{3, 1, 2, 4}
	output := sortArrayByParity(nums)
	fmt.Println(output) // Output: [2 4 3 1]

	// Example usage: The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
	nums = []int{3, 2, 3}
	fmt.Println(majorityElement(nums)) // Output: 3

	// Example usage: Unique element  in an array
	nums = []int{2, 2, 1, 1, 4, 4, 5}
	fmt.Println(singleNumber(nums)) // Output: 5
	
	// Example usage: Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4,nil}}}
	fmt.Println("Add two number of linked list:",addTwoNumbers(l1,l2))
}