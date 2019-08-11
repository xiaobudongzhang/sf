package nlogn

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestQuickSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:%v\n", arr)
	QuickSort(arr, len(arr))
	fmt.Printf("after sort:%v\n", arr)
}
