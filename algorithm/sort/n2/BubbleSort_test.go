package n2

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:%v\n", arr)
	BubbleSort(arr, len(arr))
	fmt.Printf("after sort:%v\n", arr)
}
