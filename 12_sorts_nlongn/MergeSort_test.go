package _2_sorts_nlongn

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMergeSort(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:\n", arr)
	MergeSort(arr)
	fmt.Printf("after sort:\n", arr)
}

func PrintFunc() {
	fmt.Println("--------------------------------------")
	funName, _, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}
