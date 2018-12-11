package _3_sorts_n

import (
	"testing"
	"fmt"
	"runtime"
)

func TestBucketSort(t *testing.T) {
	PrintFunc()
	arr := []int{22,5,11,41,45,26,29,10,7,8,30,27,42,43,40}
	fmt.Printf("before sort:\n", arr)
	BucketSort(arr)
	fmt.Printf("after sort:\n", arr)
}

func TestCountingSort(t *testing.T) {
	PrintFunc()
	arr := []int{22,5,11,41,45,26,29,10,7,8,5,5,30,27,42,43,40}
	fmt.Printf("before sort:\n", arr)
	CountingSort(arr)
	fmt.Printf("after sort:\n", arr)
}

func TestRadixSort(t *testing.T) {
	PrintFunc()
	arr := []int{124,132,112,107,2,127}
	fmt.Printf("before sort:\n", arr)
	RadixSort(arr)
	fmt.Printf("after sort:\n", arr)
}


func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}