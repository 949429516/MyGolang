package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newhead := reverseList(head.Next)
	post := head.Next // 递归回归  当前head=2 post赋值为3
	post.Next = head  // 3的下一个节点为2
	head.Next = nil   // 2的下一个节点为nil
	return newhead    // 返回3

}
func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reverseList(h)
}
