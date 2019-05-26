package stack_20190526___20190529

import "fmt"

/*
 *基于链表实现的栈
 */
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (This *LinkedListStack) IsEmpty() bool {
	if nil == This.topNode {
		return true
	}
	return false
}

func (This *LinkedListStack) Push(v interface{}) {
	This.topNode = &node{next: This.topNode, val: v}
}

func (This *LinkedListStack) Pop() interface{} {
	if This.IsEmpty() {
		return nil
	}

	v := This.topNode.val
	This.topNode = This.topNode.next
	return v
}

func (This *LinkedListStack) Top() interface{} {
	if This.IsEmpty() {
		return nil
	}
	return This.topNode.val
}

func (This *LinkedListStack) Flush() {
	This.topNode = nil
}

func (This *LinkedListStack) Print() {
	if This.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := This.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
