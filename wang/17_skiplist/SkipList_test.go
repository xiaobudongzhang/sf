package _7_skiplist

import (
	"testing"
	"fmt"
	"runtime"
)

func TestSkipList_Insert(t *testing.T) {
	PrintFunc()
	newList := NewSkipList()
	newList.Insert(1)
	newList.Insert(3)
	newList.Insert(4)
	newList.Insert(5)
	newList.Insert(7)
	newList.Insert(8)
	newList.Insert(9)
	newList.Insert(10)
	newList.Insert(13)
	newList.Insert(16)
	newList.Insert(17)
	newList.Insert(18)
	newList.Insert(19)
	newList.Insert(21)
	newList.Insert(23)
	newList.Insert(24)
	newList.Insert(25)
	newList.Insert(28)
	newList.Insert(30)
	newList.Insert(31)
	newList.Insert(34)
	newList.Insert(35)
	newList.Insert(36)
	newList.Insert(37)
	newList.Insert(38)
	newList.Insert(39)

	fNode := newList.Find(3)

	newList.Print()
	fmt.Printf("val:%+v---level:%+v---forward:%+v\n", fNode.val, fNode.level, fNode.forward)

	newList.Delete(34)
	newList.Print()
}


func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}