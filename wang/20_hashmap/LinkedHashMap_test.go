package _0_hashmap

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHashMap_Insert(t *testing.T) {
	PrintFunc()
	list := NewHashMap(10)
	list.Insert("3", 11)
	list.Insert("1", 12)
	list.Insert("5", 23)
	list.Insert("2", 22)
	list.Insert("7", 22)
	list.Insert("8", 22)
	list.Insert("11", 22)
	list.Insert("9", 22)
	list.Insert("13", 22)
	list.Insert("26", 22)
	list.Insert("28", 22)
	list.Insert("29", 22)
	list.Insert("31", 22)
	list.Insert("32", 22)
	list.Insert("34", 22)
	list.Insert("36", 22)
	list.Insert("38", 22)
	list.Insert("41", 22)
	list.Insert("42", 22)
	list.Insert("45", 22)
	list.Insert("48", 22)
	list.Insert("49", 4922)

	list.Print()

	list.Insert("3", 26)

	list.Print()
	//find
	findV := list.Find("49")
	fmt.Printf("find:%v\n", findV)
	//del
	fmt.Printf("-----del:%v------\n", "38")
	list.Delete("38")
	list.Print()
}
func TestHashMap_Delete(t *testing.T) {
	PrintFunc()
	list := NewHashMap(10)
	list.Insert("23", 11)
	list.Insert("31", 11)
	fmt.Printf("-----del:%v------\n", "31")
	list.Delete("31")
	list.Print()
}
func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}
