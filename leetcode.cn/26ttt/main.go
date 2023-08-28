package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treepath(node *TreeNode, s string) {

}
func binaryTreePath(root *TreeNode) []string {
	ret := []string{}
	var TreePath func(*TreeNode, string)
	TreePath = func(node *TreeNode, s string) {
		if node != nil {
			s += strconv.Itoa(node.Val)
			if node.Left == nil && node.Right == nil {
				ret = append(ret, s)
			} else {
				s += "->"
				TreePath(node.Left, s)
				TreePath(node.Right, s)
			}
		}
	}
	TreePath(root, "")
	return ret
}
func main() {
	a := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(binaryTreePath(a))
}
