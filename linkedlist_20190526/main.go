package linkedlist_20190526

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

//单链表
type LinkedList struct {
	head *ListNode
}

func (This *LinkedList) Print() {
	cur := This.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
 *单链表反转
 *时间复杂度：O(N)
 */
func (This *LinkedList) Reverse() {
	if nil == This.head || nil == This.head.next || nil == This.head.next.next {
		return
	}
	var pre *ListNode = nil
	cur := This.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	This.head.next = pre
}

/*
 *判断单链表是否有环
 */
func (This *LinkedList) HasCycle() bool {
	if nil != This.head {
		slow := This.head
		fast := This.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
 *两个有序单链表合并
 */
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}

	return l
}

/*
 *删除倒数第N个节点
 */

func (This *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == This.head || nil == This.head.next {
		return
	}

	fast := This.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}

	slow := This.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

/*
获取中间节点
*/
func (This *LinkedList) FindMiddleNode() *ListNode {
	if nil == This.head || nil == This.head.next {
		return nil
	}
	if nil == This.head.next.next {
		return This.head.next
	}

	slow, fast := This.head, This.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
