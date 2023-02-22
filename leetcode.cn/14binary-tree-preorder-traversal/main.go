package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // 前序遍历 [1 2 4 5 3 6 7]
		preorder(node.Left)
		// res = append(res, node.Val) // 中序遍历 [1 2 4 5 3 6 7]
		preorder(node.Right)
		// res = append(res, node.Val) // 后序遍历 [4 5 2 6 7 3 1]

	}
	preorder(root)
	return
}

func main() {
	n := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(preorderTraversal(n))
}
