package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(node *ListNode) *ListNode {
	// 反转链表
	var perval, curval *ListNode = nil, node
	for curval != nil {
		tmp := curval.Next

		curval.Next = perval
		perval = curval

		curval = tmp
	}
	return perval
}
func halfListnode(node *ListNode) *ListNode {
	// 获取中间位置，快指针走两步相当于2n 慢指针走一步相当于n, 2n等于nil时候n刚好为一半
	fast, slow := node, node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 获取中间位置
	halflistnode := halfListnode(head)
	// 将中间位置之后的反转
	reverselist := reverseList(halflistnode.Next)

	for reverselist != nil {
		if reverselist.Val != head.Val {
			return false
		}
		reverselist = reverselist.Next
		head = head.Next
	}
	return true
}
func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(isPalindrome(list))
}
