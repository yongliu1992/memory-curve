package array_20190523

import (
	"fmt"
	"github.com/pkg/errors"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 */

type Array struct {
		data []int
		length uint
}
const (
	ErrIndexOutOfRange = "out of index range"
	ErrFullArray = "full array"
)
//为数组初始化内存
func NewArray(capacity uint) *Array  {
	if capacity==0 {
		return nil
	}
	return &Array{
		data: make([]int,capacity,capacity),
		length:0,
	}
}

func (This *Array) Len() uint {
	return This.length
}
//判断索引是否越界
func (This *Array) IsIndexOutOfRange(index uint)bool  {
	if index>= uint(cap(This.data)) {
		return  true
	}
	return  false
}


//通过索引查找数组，索引范围[0,n-1]
func (This *Array) Find(index uint)(int,error)  {
	if This.IsIndexOutOfRange(index) {
		return 0,errors.New(ErrIndexOutOfRange)
	}
	return This.data[index],nil
}

func (This *Array) Insert(index uint,v int)error  {
	if This.Len()==uint(cap(This.data)) {
		return errors.New(ErrFullArray)
	}
	if index != This.length && This.IsIndexOutOfRange(index) {
		return errors.New(ErrIndexOutOfRange)
	}
	for i:=This.length;i>index;i-- {
		This.data[i]=This.data[i-1]
	}
	This.data[index] = v
	This.length++
	return  nil
}

//删除索引index上的值
func (This *Array) Delete(index uint) (int, error) {
	if This.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := This.data[index]
	for i := index; i < This.Len()-1; i++ {
		This.data[i] = This.data[i+1]
	}
	This.length--
	return v, nil
}

func (This *Array) Print()  {
	var format string
	for i := uint(0);i<This.Len();i++ {
		format += fmt.Sprintf("|%+v",This.data[i])
	}
	fmt.Println(format)
}