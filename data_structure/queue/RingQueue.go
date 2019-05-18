package queue

type RingQueue struct {
	head int
	tail int
	capacity int
	queue []interface{}
}

func NewRingQueue(n int) *RingQueue {
	queue := make([]interface{}, n)
	return &RingQueue{
		head:0,
		tail:0,
		capacity:n,
		queue:queue,
	}
}
func (this *RingQueue)EnQueue(ele interface{}) bool  {
	if (this.tail + 1) % this.capacity == this.head {
		return false
	}
	this.queue[this.tail] =  ele
	this.tail = (this.tail + 1) % this.capacity
	return true
}

func (this *RingQueue)DeQueue() interface{} {
	if this.tail == this.head {
		return nil
	}
	ele := this.queue[this.head]
	this.head = (this.head + 1) % this.capacity
	return ele
}

