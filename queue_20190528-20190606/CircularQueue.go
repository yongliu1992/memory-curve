package queue_20190528_20190606

import "fmt"

type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

/*
 *栈空条件：head==tail为true
 */
func (This *CircularQueue) IsEmpty() bool {
	if This.head == This.tail {
		return true
	}
	return false
}

/*
 *栈满条件：(tail+1)%capacity==head为true
 */

func (This *CircularQueue) IsFull() bool {
	if This.head == (This.tail+1)%This.capacity {
		return true
	}
	return false
}

func (This *CircularQueue) EnQueue(v interface{}) bool {
	if This.IsEmpty() {
		return false
	}
	This.q[This.tail] = v
	This.tail = (This.tail + 1) % This.capacity
	return true
}

func (This *CircularQueue) DeQueue() interface{} {
	if This.IsEmpty() {
		return nil
	}
	v := This.q[This.head]
	This.head = (This.head + 1) % This.capacity
	return v
}

func (This *CircularQueue) String() string {
	if This.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = This.head
	for true {
		result += fmt.Sprintf("<-%+v", This.q[i])
		i = (i + 1) % This.capacity
		if i == This.tail {
			break
		}
	}
	result += "<-tail"
	return result

}
