package queue_20190528_20190606

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (This *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if nil == This.tail {
		This.tail = node
		This.head = node
	} else {
		This.tail.next = node
		This.tail = node
	}
	This.length++
}

func (This *LinkedListQueue) DeQueue() interface{} {
	if This.head == nil {
		return nil
	}
	v := This.head.val
	This.head = This.head.next
	This.length--
	return v
}

func (This *LinkedListQueue) String() string {
	if This.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := This.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
