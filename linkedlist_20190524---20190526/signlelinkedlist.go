package linkedlist_20190524___20190526

import "fmt"

/**
 * 单链表基本操作
 */

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (This *ListNode) GetNextNode() *ListNode {
	return This.next()
}

func (This *ListNode) GetValue() interface{} {
	return This.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (This *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := newNode.next
	p.next = newNode
	newNode.next = oldNext
	This.length++
	return true
}
func (This *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == This.head {
		return false
	}
	cur := This.head.next
	pre := This.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	This.length++
	return true
}

//在链表头部插入节点
func (This *LinkedList) InsertToHead(v interface{}) bool {
	return This.InsertAfter(This.head, v)
}

//在链表尾部插入节点
func (This *LinkedList) InsertToTail(v interface{}) bool {
	cur := This.head
	for nil != cur.next {
		cur = cur.next
	}
	return This.InsertAfter(cur, v)
}

//通过索引查找节点
func (This *LinkedList) FindByIndex(index uint) *ListNode {
	if index > This.length {
		return nil
	}
	cur := This.head
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (This LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := This.head.next
	pre := This.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	This.length--
	return true
}

func (This *LinkedList) Print() {
	cur := This.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
