package _9_queue


import "fmt"

type CircularQueue struct {
	queue []interface{}
	capacity int
	head int
	tail int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		queue:make([]interface{}, n, n),
		capacity:n,
		head:0,
		tail:0,
	}
}

func (this *CircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *CircularQueue) IsFull() bool {
	return (this.tail + 1) % this.capacity == this.head
}

func (this *CircularQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return true
}

func (this *CircularQueue) DeQueue() interface{}{
	if this.IsEmpty() {
		return nil
	}
	hval := this.queue[this.head]
	this.head = (this.head + 1) % this.capacity
	return hval
}

func (this *CircularQueue) String() string {
	if this.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = this.head
	for {
		result += fmt.Sprintf("<-%+v", this.queue[i])
		i = (i + 1) % this.capacity
		if i == this.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
