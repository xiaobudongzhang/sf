package n

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestRadixSort(t *testing.T) {
	lib.PrintFunc()
	arr := []int{18800345665, 132, 17703771999, 15371842777, 18016533288, 127}
	fmt.Printf("before sort:%v\n", arr)
	RadixSort(arr)
	fmt.Printf("after sort:%v\n", arr)
}
