/*
设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
你的实现应该支持如下操作：
MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。

示例：
MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
circularQueue.enQueue(1);  // 返回 true
circularQueue.enQueue(2);  // 返回 true
circularQueue.enQueue(3);  // 返回 true
circularQueue.enQueue(4);  // 返回 false，队列已满
circularQueue.Rear();  // 返回 3
circularQueue.isFull();  // 返回 true
circularQueue.deQueue();  // 返回 true
circularQueue.enQueue(4);  // 返回 true
circularQueue.Rear();  // 返回 4
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/design-circular-queue
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type MyCircularQueue struct {
	start, end    *ListNode
	maxSize, size int
}

func Constructor(k int) *MyCircularQueue {
	return &MyCircularQueue{maxSize: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &ListNode{Val: value}
	if this.start == nil {
		this.start = node
		this.end = node
	} else {
		this.end.Next = node
		this.end = node
	}
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = this.start.Next
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.start.Val
}
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.end.Val
}
func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}
func (this *MyCircularQueue) IsFull() bool {
	if this.size == this.maxSize {
		return true
	}
	return false
}
func main() {
	m := Constructor(3)
	m.EnQueue(1)
	m.EnQueue(2)
	m.EnQueue(3)
	fmt.Println(m.IsEmpty())
	fmt.Println(m.Front())
	fmt.Println(m.Rear())
}
