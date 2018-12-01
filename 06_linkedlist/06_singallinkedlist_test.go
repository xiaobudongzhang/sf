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
	fmt.Printf("xv: %v", cur.GetValue())
}

func PrintFunc()  {
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}