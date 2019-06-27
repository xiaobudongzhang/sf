package heap

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(10)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(7)
	heap.Insert(2)
	heap.Insert(1)
	heap.Print()
}

func TestHeap_RemoveMax(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(10)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(7)
	heap.Insert(2)
	heap.Insert(1)
	heap.Print()

	fmt.Printf("--------------------------------------------------------------------------------------------------------------\n")
	heap.RemoveMax()
	heap.Print()
}

func TestHeap_Sort(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(100)
	sortData := []int{0, 2,6,1,7,3,8,5} //todo 需要从下标1开始
	heap.Sort(sortData)
	fmt.Printf("data:%v\n", sortData)
}
