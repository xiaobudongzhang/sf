package queue

import "fmt"

type Node struct {
	next *Node
	ele interface{}
}

type LinkedQueue struct {
	head *Node
	tail *Node
}

func NewNode(ele interface{}) *Node  {
	return &Node{
		ele:ele,
		next:nil,
	}
}
func NewLinkedQueue() *LinkedQueue  {
	return &LinkedQueue{
		head:nil,
		tail:nil,
	}
}

func (this *LinkedQueue) Enqueue(ele interface{}) bool {
	node := NewNode(ele)
	if this.tail != nil {
		this.tail.next = node
	}
	this.tail = node

	if this.head == nil {
		this.head = node
	}
	return true
}

func (this *LinkedQueue) Dequeue() interface{} {
	if this.head == nil {
		return nil
	}
	ele := this.head.ele
	this.head = this.head.next
	if this.head == nil {//if not do this, this tail pointer to a node
		this.tail = nil
	}
	return ele
}

func (this *LinkedQueue) Print()  {
	fmt.Printf("%v\n", this.tail)
	cur := this.head
	for cur != nil  {
		fmt.Printf("-->%v", cur.ele)
		cur = cur.next
	}
	fmt.Println()
}