package _8_heap

import (
	"lib"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(100)
	heap.Insert(1)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(3)
	heap.Insert(9)

	heap.Print()
}

func TestHeap_RemoveMax(t *testing.T) {
	lib.PrintFunc()
	heap := NewHeap(100)
	heap.Insert(1)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(3)
	heap.Insert(9)

	heap.RemoveMax()
	heap.Print()
}
