package _11_sorts

import (
	"testing"
	"fmt"
	"runtime"
)

func TestBubbleSort(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:\n", arr)
	BubbleSort(arr, len(arr))
	fmt.Printf("after sort:\n", arr)
}

func TestInsertSort(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:\n", arr)
	InsertSort(arr, len(arr))
	fmt.Printf("after sort:\n", arr)
}

func TestSelectionSort(t *testing.T) {
	PrintFunc()
	arr := []int{3, 5, 4, 1, 2, 6}
	fmt.Printf("before sort:\n", arr)
	SelectionSort(arr, len(arr))
	fmt.Printf("after sort:\n", arr)
}
func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}
