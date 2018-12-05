package _7_stack

import "fmt"

type Browser struct {
	prev *ArrayStack
	next *ArrayStack
}

func NewBrowser() *Browser {
	return &Browser{
		prev:NewArrayStack(3),
		next:NewArrayStack(3),
	}
}

func (this *Browser) Prev() {
	if this.next.IsEmpty(){
		return
	}
	this.prev.Push(this.next.Pop())
	fmt.Printf("Prev :%v\n", this.prev.Top())
}

func (this *Browser) Next() {
	if this.prev.IsEmpty() {
		return
	}
	this.next.Push(this.prev.Pop())
	fmt.Printf("Next :%v\n", this.next.Top())
}

func (this *Browser) PushPrev(v interface{}) {
	this.prev.Push(v)
}

func (this *Browser) Open(v interface{}) {
	this.prev.Push(v)
	this.next.Flush()
	fmt.Printf("Open :%v\n", v)
}