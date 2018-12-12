package _5_binarysearch

import (
	"testing"
	"fmt"
	"runtime"
)

func TestBinarySearch(t *testing.T) {
	PrintFunc()
	a := []int{1,2,3,4,5,6,7,8,9}
	index := BinarySearch(a, 4)
	fmt.Printf("index:%v\n", index)
}

func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}