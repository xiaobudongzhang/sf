package _6_linkedlist

import (
	"fmt"
)

type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil,v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(nil), 0}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)

	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext

	this.length++
	return true
}

func (this *LinkedList) FindPrev(p *ListNode) *ListNode {
	prev := this.head
	for   {
		if prev.next == p {
			break
		}
		prev = prev.next
	}
	return prev
}
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool  {
	if nil == p || this.head == p {
		return false
	}
	newNode := NewListNode(v)
	//find prev node of p
	prev := this.FindPrev(p)
	newNode.next = p
	prev.next = newNode
	this.length++
	return true
}

func (this *LinkedList) InsertToHead(v interface{}) bool {
	this.InsertAfter(this.head, v)
	return true
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	this.InsertAfter(cur, v)
	return true
}


func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}

	return cur
}

func (this *LinkedList) FindByVal(val interface{}) *ListNode {
	if val == nil {
		return nil
	}
	cur := this.head.next
	for cur != nil {
		if cur.value == val {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p || p == this.head {
		return false
	}
	//find p
	prev := this.FindPrev(p)
	prev.next = p.next
	p = nil
	this.length--
	return true
}

func (this *LinkedList) DeleteLastNode() bool {
	if nil == this.head.next {
		return false
	}
	prev := this.head
	cur := this.head.next

	for nil != cur {
		cur = cur.next
		prev = prev.next
	}
	cur = nil
	this.length--
	return true
}

func (this *LinkedList) Reverse(start *ListNode) bool {
	if this.length < 4 {
		return true
	}
	var newLinkHead *ListNode = nil
	cur := start.next

	for nil != cur  {
		oldNext := cur.next
		cur.next = newLinkHead
		newLinkHead = cur
		cur = oldNext
	}
	start.next = newLinkHead
	return true
}

func (this *LinkedList) HasCycle() bool  {
	slow := this.head
	fast := this.head
	for nil != slow.next && nil != fast.next.next {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func (this *LinkedList) MergeSortedList(that *LinkedList) *LinkedList  {
	linkedList := NewLinkedList()

	thisCur := this.head.next
	thatCur := that.head.next
	newCur := linkedList.head

	for nil != thisCur && nil != thatCur {
		if thisCur.value.(int) < thatCur.value.(int) {
			newCur.next = thisCur
			thisCur = thisCur.next
			newCur = newCur.next
		} else if thisCur.value.(int) > thatCur.value.(int) {
			newCur.next = thatCur
			thatCur = thatCur.next
			newCur = newCur.next
		} else {
			oldThisCur := thisCur
			newCur.next = oldThisCur

			thisCur = thisCur.next

			newCur.next.next = thatCur
			thatCur = thatCur.next
			newCur = newCur.next.next
		}
	}
	if thisCur != nil {
		newCur.next = thisCur
	}
	if thatCur != nil {
		newCur.next = thatCur
	}
	linkedList.length = this.length + that.length
	return linkedList
}

func (this *LinkedList) DeleteBottomN(n int) bool{
	if this.head.next == nil {
		return true
	}
	fastP := this.head.next
	for i := 0; i<= n && fastP != nil; i++ {
		fastP = fastP.next
	}

	slowP := this.head.next
	for nil != fastP {
		slowP = slowP.next
		fastP = fastP.next
	}
	if slowP == nil {
		return true
	}
	next := slowP.next
	slowP.next = next.next
	next = nil
	return true
}

func (this *LinkedList) FindMiddleNode() *ListNode{
	if this.head.next == nil {
		return NewListNode(nil)
	}
	fast := this.head.next
	slow := this.head.next

	for nil != fast && nil != fast.next {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func (this *LinkedList) Print()  {

	cur := this.head.next
	fmt.Printf("length:%d\n", this.length)
	for i := uint(0); i <= this.length ; i++ {
		if (cur == nil){
			break
		}
		fmt.Printf("value:%+v\n", cur.GetValue())
		cur = cur.next
	}
}

func (this *LinkedList) RemoveNthFromEnd(n int) *ListNode {
	//快慢指针来一趟遍历,慢指针指向被删除的元素的前一个元素
	fmt.Printf("%v\n", this.head)
	slow := this.head
	p :=  this.head
	i := 0


	for p.next != nil {
		break
		if i < n {
			i++
		}else{
			slow = slow.next
		}
		p = p.next

	}

	slow.next = slow.next.next
	this.length--
	return this.head
}