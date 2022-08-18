/*
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
广度优先 深度优先 生成二叉树
示例 1：


输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15
示例 2：
输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：19
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/deepest-leaves-sum
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 生成二叉树
func tree(sli []interface{}) (retTree *TreeNode) {
	if len(sli) == 0 {
		return nil
	}
	que := []*TreeNode{} // 作为队列 先进先出
	full_left := true    // 节点的左侧没有值为true则可以给左侧赋值
	// 1. 从切片中获取值
	for _, v := range sli {
		if v != nil {
			vint := v.(int)
			// 2.生成根节点,如果队列里没有值则定义根节点
			if len(que) == 0 {
				retTree = &TreeNode{Val: vint}
				que = append(que, retTree)
			} else {
				// 3.将片中的值生成一个节点，并且加入到队列
				tree := &TreeNode{Val: vint}
				que = append(que, tree)
				if full_left {
					// 4.左侧没有值则赋值给队列第一个元素的左侧,并将左侧判断参数改为false
					que[0].Left = tree
					full_left = false
				} else {
					// 赋值给右侧并从队列删除
					que[0].Right = tree
					full_left = true
					que = que[1:]
				}
			}
		} else {
			if len(sli) == 0 {
				continue
			} else {
				if full_left {
					full_left = false
				} else {
					full_left = true
					que = que[1:]
				}
			}
		}
	}
	return
}

// 广度优先
func deepestLeavesSum(root *TreeNode) (ret int) {
	// 1.创建队列存放值
	que := []*TreeNode{root}
	// 2.遍历队列，直到队列为空
	for len(que) > 0 {
		ret = 0 // 每层的数值总和
		// 3.计算当前层的总和，并且将该层的下一节加入到队列
		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			ret += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return
}

// 深度优先
func ddeepestLeavesSum(root *TreeNode) (ret int) {
	maxdeep := 0                 // 记录当前最大层数
	var dfs func(*TreeNode, int) // 定义函数类型，函数内调用函数
	dfs = func(node *TreeNode, level int) {
		if node == nil { // 空的节点直接返回
			return
		}
		if level > maxdeep { // 如果当前层数大于最大层数，则说明还有下一层，1.更新最大层数 2.更新最大值,从开始算起
			maxdeep = level
			ret = node.Val
		} else if level == maxdeep { // 如果层数相等，则说明是同一层则自加当前值
			ret += node.Val
		}
		dfs(node.Left, level+1) //递归调用
		dfs(node.Right, level+1)
	}
	// 从根节点开始，设置为0层数
	dfs(root, 0)
	return
}
func main() {
	a := tree([]interface{}{1, 2, 3, 4, 5, nil, 6, 7, nil, nil, nil, nil, 8})
	fmt.Println(deepestLeavesSum(a))
	fmt.Println(ddeepestLeavesSum(a))
}
