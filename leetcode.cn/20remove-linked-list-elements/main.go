package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{Next: head}
	for newh := h; newh.Next != nil; {
		if newh.Next.Val == val {
			newh.Next = newh.Next.Next
		} else {
			newh = newh.Next
		}
	}
	return h.Next
}
func main() {
	h1 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{6, &ListNode{4, &ListNode{6, nil}}}}}}
	removeElements(&h1, 6)
}
