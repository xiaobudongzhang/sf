package _6_linkedlist

import (
	"fmt"
	"runtime"
	"testing"
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
	if linkedList.length%2 != 0 {
		step++
	}
	midP := linkedList.FindByIndex(step)

	linkedList.Reverse(midP)
	linkedList.Print()
}
func TestLinkedList_HasCycle(t *testing.T) {
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

	lastp := linkedList.FindByIndex(linkedList.length - 1)
	lastp_3 := linkedList.FindByIndex(linkedList.length - 3)
	lastp.next = lastp_3
	fmt.Printf("%v", linkedList.HasCycle())
}

func TestLinkedList_MergeSortedList(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()
	linkedList2 := NewLinkedList()

	linkedList.InsertToTail(1)
	linkedList.InsertToTail(7)
	linkedList.InsertToTail(9)

	linkedList2.InsertToTail(1)
	linkedList2.InsertToTail(3)
	linkedList2.InsertToTail(4)
	linkedList2.InsertToTail(8)
	linkedList2.InsertToTail(10)
	linkedList2.InsertToTail(15)

	linkedList3 := linkedList2.MergeSortedList(linkedList)
	linkedList3.Print()
}

func TestLinkedList_DeleteBottomN(t *testing.T) {
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
	linkedList.InsertToTail("newNode10")

	linkedList.DeleteBottomN(-12)
	linkedList.Print()
}

func TestLinkedList_FindMiddleNode(t *testing.T) {
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
	linkedList.InsertToTail("newNode10")

	midP := linkedList.FindMiddleNode()
	fmt.Printf("%v", midP.value)
}
func PrintFunc() {
	fmt.Println("--------------------------------------")
	funName, _, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}

func TestLinkedList_RemoveNthFromEnd(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()
	linkedList.InsertToTail(1)

	linkedList.RemoveNthFromEnd(1)
	linkedList.Print()
}
