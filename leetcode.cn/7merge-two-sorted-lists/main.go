package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// if list1 == nil {
	// 	return list2
	// }
	// if list2 == nil {
	// 	return list1
	// }
	// if list1.Val > list2.Val {
	// 	list1, list2 = list2, list1
	// }
	// list1.Next = mergeTwoLists(list1.Next, list2)
	// return list1
	retn := &ListNode{-1, nil}
	ret := retn
	for {
		if list1 == nil {
			ret.Next = list2
			return retn.Next
		}
		if list2 == nil {
			ret.Next = list1
			return retn.Next
		}
		if list1.Val >= list2.Val {
			ret.Next = list2
			list2 = list2.Next
		} else {
			ret.Next = list1
			list1 = list1.Next
		}
		ret = ret.Next
	}
}

func main() {
	a := ListNode{
		1, &ListNode{
			2, &ListNode{
				3, nil,
			},
		},
	}
	b := ListNode{
		2, &ListNode{
			4, nil,
		},
	}
	c := mergeTwoLists(&a, &b)
	for {
		fmt.Println(c.Val)
		c = c.Next
		if c.Next == nil {
			fmt.Println(c.Val)
			break
		}
	}
}
