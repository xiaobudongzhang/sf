package n

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestCountingSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{22, 5, 11, 41, 45, 26, 29, 10, 7, 8, 5, 5, 30, 27, 42, 43, 40}
	fmt.Printf("before sort:%v\n", arr)
	CountingSort(arr)
	fmt.Printf("after sort:%v\n", arr)
}
