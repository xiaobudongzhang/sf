package _5_binarysearch

import (
	"testing"
	"fmt"
)

func TestFirstEqual(t *testing.T) {
	PrintFunc()
	a := []int{1,7,8,8,8,8,8,8}
	index := FirstEqual(a)
	fmt.Printf("index:%v\n", index)
}
