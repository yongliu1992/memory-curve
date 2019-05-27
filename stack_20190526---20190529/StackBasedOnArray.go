package stack_20190526___20190529

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (This *ArrayStack) IsEmpty() bool {
	if This.top < 0 {
		return true
	}
	return false
}

func (This *ArrayStack) Push(v interface{}) {
	if This.top < 0 {
		This.top = 0
	} else {
		This.top += 1
	}

	if This.top > len(This.data)-1 {
		This.data = append(This.data, v)
	} else {
		This.data[This.top] = v
	}
}

func (This *ArrayStack) Pop() interface{} {
	if This.IsEmpty() {
		return nil
	}
	v := This.data[This.top]
	This.top -= 1
	return v
}

func (This *ArrayStack) Top() interface{} {
	if This.IsEmpty() {
		return nil
	}
	return This.data[This.top]
}

func (This *ArrayStack) Flush() {
	This.top = -1
}

func (This *ArrayStack) Print() {
	if This.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := This.top; i >= 0; i-- {
			fmt.Println(This.data[i])
		}
	}
}
