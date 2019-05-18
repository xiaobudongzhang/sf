package queue

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestLinkedQueue_Enqueue(t *testing.T) {
	lib.PrintFunc()


	queue := NewLinkedQueue()
	for i:=1; i< 5; i++  {
		queue.Enqueue(i)
	}
	queue.Print()

	for i:=1; i< 5; i++  {
		ele := queue.Dequeue()
		fmt.Printf("dequeue:%v\n", ele)
	}


	queue.Print()


	for i:=11; i< 15; i++  {
		queue.Enqueue(i)
	}
	queue.Print()
}
