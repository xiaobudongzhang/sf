package _3_tree

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNewTree(t *testing.T) {
	PrintFunc()
	datas := []interface{}{"","A","B","C","D","E","F","G","H","I","J"}
	tree := NewTree(datas)
	tree.Print(1)
}
func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}
