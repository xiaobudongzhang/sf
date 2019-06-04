package queue

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestRingQueue_EnQueue(t *testing.T) {
	lib.PrintFunc()

	queue := NewRingQueue(4)
	for i := 1; i < 5; i++ {
		req := i
		res := queue.EnQueue(req)
		fmt.Printf("%v:%v:%v\n", req, res, queue)
	}

	fmt.Println("--------------------------DeQueue---------------------------")
	res := queue.DeQueue()
	fmt.Printf("%v:%v\n", res, queue)

	fmt.Println("--------------------------EnQueue---------------------------")
	req := 4
	res = queue.EnQueue(req)
	fmt.Printf("%v:%v:%v\n", req, res, queue)

}
