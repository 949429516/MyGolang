/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。

实现 CBTInserter 类:

CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态，并返回插入节点 TreeNode 的父节点的值；
CBTInserter.get_root() 将返回树的头节点。
示例 1：

输入
["CBTInserter", "insert", "insert", "get_root"]
[[[1, 2]], [3], [4], []]
输出
[null, 1, 2, [1, 2, 3, 4]]
解释
CBTInserter cBTInserter = new CBTInserter([1, 2]);
cBTInserter.insert(3);  // 返回 1
cBTInserter.insert(4);  // 返回 2
cBTInserter.get_root(); // 返回 [1, 2, 3, 4]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/complete-binary-tree-inserter
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type CBTInserter struct {
	// 存放根节点
	root *TreeNode
	// 存放可以添加子结点的候选节点,该节点右侧或左右可能没有值
	candidate []*TreeNode
}

// 返回结构体
func Constructor(root *TreeNode) CBTInserter {
	ro := []*TreeNode{root}
	// 存放没有子结点或者有一个子结点的
	candidate := []*TreeNode{}
	for len(ro) > 0 {
		node := ro[0]
		ro = ro[1:]
		// 判断左右是否有左右节点，1.如果有则将加入ro中(可能其子结点还有子结点)2.如果没有则加入candidate
		if node.Left != nil {
			ro = append(ro, node.Left)
		}
		if node.Right != nil {
			ro = append(ro, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

// 将val加入二叉树
func (this *CBTInserter) Insert(val int) int {
	// 创建child结构体
	child := &TreeNode{Val: val}
	// 将child加入第一个candidate中的值
	node := this.candidate[0]
	if node.Left == nil {
		node.Left = child
	} else {
		// 如果右侧加入，则说明该节点都有子结点了，则在candidate删除掉
		node.Right = child
		this.candidate = this.candidate[1:]
	}
	// child加入candidate
	this.candidate = append(this.candidate, child)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
