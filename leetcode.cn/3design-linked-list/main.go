package main

/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，
则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。
如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

示例：
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/design-linked-list
*/
import "fmt"

type node struct {
	val        int
	next, prev *node
}
type MyLinkedList struct {
	size       int
	head, tail *node
}

// 初始化双向链表
func Constructor() MyLinkedList {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{0, head, tail}
}

// 获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	// 双向链表，如果在前半段则从前寻找，后半段从后寻找
	var current *node
	if index+1 < this.size-index {
		current = this.head
		for i := 0; i <= index; i++ {
			current = current.next
		}
	} else {
		current = this.tail
		for i := 0; i < this.size-index; i++ {
			current = current.prev
		}
	}
	return current.val
}

// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

// 在链表中的第index个节点之前添加值为val的节点。
// 1.如果 index 等于链表的长度，则该节点将附加到链表的末尾
// 2.如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	var before, after *node
	// 找到index的插入位置，获取到该位置之前和之后的节点
	if index < this.size-index {
		before = this.head
		for i := 0; i < index; i++ {
			before = before.next
		}
		after = before.next
	} else {
		after = this.tail
		for i := 0; i < this.size-index; i++ {
			after = after.prev
		}
		before = after.prev
	}
	// 将新node与其之前和之后的连接起来
	this.size++
	addnode := &node{val, after, before}
	before.next = addnode
	after.prev = addnode
}

// 如果索引 index 有效，则删除链表中的第 index 个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	// 找到index节点，找到其前、后节点
	var before, after *node
	if index < this.size-index {
		before = this.head
		for i := 0; i < index; i++ {
			before = before.next
		}
		after = before.next.next
	} else {
		after = this.tail
		for i := 0; i < this.size-index-1; i++ {
			after = after.prev
		}
		before = after.prev.prev
	}
	// 跳过index节点，将他两端的节点连接
	this.size--
	before.next = after
	after.prev = before
}
func main() {
	L := Constructor()
	L.AddAtHead(1)
	L.AddAtTail(3)
	L.AddAtIndex(1, 2)
	fmt.Println(L.Get(1))
	L.DeleteAtIndex(1)
	fmt.Println(L.Get(1))

}
