package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Rigth *TreeNode
}

/*
stack先进后出原则
*/
func inorderTraversal(root *TreeNode) (res []int) {
	// 栈，先进后出
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 入栈,左树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Rigth
	}
	return
}
func main() {
	a := TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(inorderTraversal(&a))
}
