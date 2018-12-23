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

