package _8_heap

import (
	"fmt"
	"lib"
	"testing"
)

func TestHeap_Sort(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(200)

	fmt.Printf("--init--\n")
	heap.Init([]int{7, 16, 20, 13, 4, 1, 19, 5, 8})
	heap.Print()

	fmt.Printf("--BuildHeap--\n")
	heap.BuildHeap()
	heap.Print()

	fmt.Printf("--Sort--\n")
	heap.Sort()
	heap.Print()

	fmt.Printf("data:%v\n", heap.data)
}
