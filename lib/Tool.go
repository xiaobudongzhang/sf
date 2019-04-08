package lib

import (
	"fmt"
	"runtime"
)

func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}

func Hash33(k interface{}, size int) int{
	//time33
	hash := 5381
	for _,v := range k.(string) {
		hash += (hash << 5) + int(v)
	}
	return (hash & 0x7FFFFFFF) % size
}


func Min(a []int)  int{
	min := a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	return min
}

func Max(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func MinInt(a,b int) int {
	if a > b {
		return b
	}
	return a
}
