package _8_stack

type ArrayStack struct {
	stack []interface{}
	top   int
}

func NewArrayStack(n uint) *ArrayStack {
	return &ArrayStack{
		stack: make([]interface{}, n, n),
		top:   -1,
	}
}

func (this *ArrayStack) Push(val interface{}) {
	if this.top+1 >= len(this.stack) {
		return
	}
	this.top++
	this.stack[this.top] = val
}

func (this *ArrayStack) Pop() interface{} {
	if this.top == -1 {
		return nil
	}
	v := this.stack[this.top]
	this.top--
	return v
}

func (this *ArrayStack) IsEmpty() bool {
	return this.top == -1
}

func (this *ArrayStack) Top() interface{} {
	return this.stack[this.top]
}

func (this *ArrayStack) Flush() {
	this.top = -1
}
