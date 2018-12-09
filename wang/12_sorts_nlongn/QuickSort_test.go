package _2_sorts_nlongn

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:\n", arr)
	QuickSort(arr)
	fmt.Printf("after sort:\n", arr)
}
