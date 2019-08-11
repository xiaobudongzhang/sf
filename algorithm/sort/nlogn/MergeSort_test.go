package nlogn

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestMergeSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:%v\n", arr)
	MergeSort(arr, len(arr))
	fmt.Printf("after sort:%v\n", arr)
}
