package _9_queue

import "fmt"

type ListNode struct {
	val interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

func NewListNode(v interface{}) *ListNode  {
	return &ListNode{
		val:v,
		next:nil,
	}
}
func NewLinkedListQueue() *LinkedListQueue  {
	return &LinkedListQueue{
		head:nil,
		tail:nil,
		length:0,
	}
}

func (this *LinkedListQueue) EnQueue(v interface{})  {
	node := NewListNode(v)
	if nil == this.tail {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{}  {
	if nil == this.head {
		return nil
	}
	oldNode := this.head
	val := oldNode.val
	this.head = this.head.next
	this.length--
	if this.head == nil {
		this.tail = nil
	}
	fmt.Printf("oldNode:%v\n", oldNode)
	return val
}


func (this *LinkedListQueue) GetTail() *ListNode  {
	return this.tail
}
func (this *LinkedListQueue) SetTail(v *ListNode )  {
	this.tail = v
}
func (this *LinkedListQueue) String() string {
	fmt.Printf("tail:%v\n", this.tail)
	if this.head == nil {
		return "empty queue"
	}
	result := "head<-"

	for cur := this.head; cur != nil; cur=cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}

	result += "<-tail"
	return result
}