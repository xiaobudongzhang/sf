package heap

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

type  VS struct {
	K int
	V int
}
var comparator = func(a interface{}, b interface{}) int {
		if a.(VS).V > b.(VS).V {
			return 1
		} else {
			return -1
		}
}

func TestHeap_Insert(t *testing.T) {
	lib.PrintFunc()

	heap := NewHeap(10, comparator)
	heap.Insert(VS{K:4,V:15})
	heap.Insert(VS{K:2,V:13})
	heap.Insert(VS{K:5,V:24})

	heap.Print()

	heap.RemoveMin()
	heap.Update(VS{K:2,V:25},VS{K:2,V:13})
	heap.Print()

}