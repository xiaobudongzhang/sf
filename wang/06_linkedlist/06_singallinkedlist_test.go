package _6_linkedlist

import (
	"testing"
	"fmt"
	"runtime"
)

func TestLinkedList_InsertToHead(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()
	linkedList.InsertToHead("InsertToHead")
	linkedList.Print()
}

func TestLinkedList_InsertToTail(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()
	linkedList.InsertToTail("InsertToTail")
	linkedList.Print()
}

func TestLinkedList_InsertBefore(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()

	linkedList.InsertToTail("newNode1")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode4")
	linkedList.InsertBefore(linkedList.head.next.next, "InsertBefore")
	linkedList.Print()
}

func TestLinkedList_DeleteNode(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()

	linkedList.InsertToTail("newNode1")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode4")
	linkedList.DeleteNode(linkedList.head.next.next)
	linkedList.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()

	linkedList.InsertToTail("newNode1")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode4")
	cur := linkedList.FindByIndex(2)
	fmt.Printf("xv: %v\n", cur.GetValue())
}

func TestLinkedList_Reverse(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()

	linkedList.InsertToTail("newNode1")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode4")
	linkedList.InsertToTail("newNode5")
	linkedList.InsertToTail("newNode6")
	linkedList.InsertToTail("newNode7")
	linkedList.InsertToTail("newNode8")
	linkedList.InsertToTail("newNode9")

	step := linkedList.length / 2
	if linkedList.length % 2 != 0{
		step++
	}
	midP := linkedList.FindByIndex(step)

	linkedList.Reverse(midP)
	linkedList.Print()
}
func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}