package _5_binarysearch

import (
	"fmt"
	"testing"
)

func TestFirstEqual(t *testing.T) {
	PrintFunc()
	a := []int{1, 7, 8, 8, 8, 8, 8, 8}
	index := FirstEqual(a)
	fmt.Printf("index:%v\n", index)
}

func TestLastEqual(t *testing.T) {
	PrintFunc()
	a := []int{1, 7, 8, 8, 8, 8, 9, 11}
	index := LastQqual(a)
	fmt.Printf("index:%v\n", index)
}

func TestFirstGtOrEqual(t *testing.T) {
	PrintFunc()
	a := []int{1, 7, 8, 8, 8, 8, 9, 11}
	index := FirstGtOrEqual(a)
	fmt.Printf("index:%v\n", index)
}

func TestLastLessOrEqual(t *testing.T) {
	PrintFunc()
	a := []int{1, 7, 8, 8, 8, 8, 9, 11}
	index := LastLessOrEqual(a)
	fmt.Printf("index:%v\n", index)
}
