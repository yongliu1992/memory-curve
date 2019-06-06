package queue_20190528_20190606

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (This *ArrayQueue) EnQueue(v interface{}) bool {
	if This.tail == This.capacity {
		return false
	}
	This.q[This.tail] = v
	This.tail++
	return true
}

func (This *ArrayQueue) DelQueue() interface{} {
	if This.head == This.tail {
		return nil
	}
	v := This.q[This.head]
	This.head++
	return v
}

func (This *ArrayQueue) String() string {
	if This.head == This.tail {
		return "empty queue"
	}
	result := "head"
	for i := This.head; i <= This.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", This.q[i])
	}
	result += "<-tail"
	return result
}
