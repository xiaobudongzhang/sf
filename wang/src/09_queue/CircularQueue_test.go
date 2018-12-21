package _9_queue

import (
	"testing"
	"fmt"
	"runtime"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	PrintFunc()
	queue := NewCircularQueue(5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.EnQueue(6)

	t.Log(queue)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	PrintFunc()
	q := NewCircularQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	t.Log(q.DeQueue())
	t.Log(q)
	q.EnQueue(5)
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
}

func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}