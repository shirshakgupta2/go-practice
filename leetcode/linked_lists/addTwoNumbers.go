package main

import "strconv"

	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		var s1, s2 string

		for l1 != nil {
			s1 = strconv.Itoa(l1.Val)+s1
			l1 = l1.Next
		}
		for l2 != nil {
			s2 = strconv.Itoa(l2.Val)+s2
			l2 = l2.Next
		}
		i1, err1 := strconv.Atoi(s1)
		i2, err2 := strconv.Atoi(s2)
		if err1 != nil || err2 != nil {
			return nil
		}
		sum := i1 + i2
		return createLinkedListFromInt(sum)
	}

	func createLinkedListFromInt(sum int) *ListNode {
		str := strconv.Itoa(sum)
		dummy := &ListNode{}   
		curr := dummy

		for i:=len(str)-1;i>=0;i-- {
			val, _ := strconv.Atoi(string(str[i]))
			curr.Next = &ListNode{Val: val}
			curr = curr.Next
		}
		return dummy
	}