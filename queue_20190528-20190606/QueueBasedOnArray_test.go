package queue_20190528_20190606

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}

func TestArrayQueue_DelQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	q.DelQueue()
	t.Log(q)
	q.DelQueue()
	t.Log(q)
	q.DelQueue()
	t.Log(q)
	q.DelQueue()
	t.Log(q)
	q.DelQueue()
	t.Log(q)
}
