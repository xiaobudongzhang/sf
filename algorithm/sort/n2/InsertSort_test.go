package n2

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestInsertSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{4, 5, 6, 1, 3, 2}
	fmt.Printf("before sort:%v\n", arr)
	InsertSort(arr, len(arr))
	fmt.Printf("after sort:%v\n", arr)
}
