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