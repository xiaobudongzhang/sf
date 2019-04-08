package _9_queue

import (
	"testing"
)

func TestLinkedListQueue_EnQueue(t *testing.T) {
	PrintFunc()
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	t.Log(q)
}


func TestLinkedListQueue_DeQueue(t *testing.T) {
	PrintFunc()
	q := NewLinkedListQueue()
	q.EnQueue(1)
	tmp := q.GetTail()
	tmp.val =2
	t.Log(tmp)
	t.Log(q.GetTail())

	tmp = nil
	t.Log(tmp)
	q.SetTail( nil)
	t.Log(q.GetTail())
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	t.Log(q)
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	t.Log(q)
	t.Log(q.DeQueue())
	t.Log(q)
	q.EnQueue(6)
	t.Log(q)
}
