package _2_sorts_nlongn

import (
	"testing"
	"fmt"
)

func TestFindLastK(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	lastK := FindLastK(arr, 3)
	fmt.Printf("lastK:%v\n", lastK)
}