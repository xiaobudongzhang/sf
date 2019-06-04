package _6_linkedlist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	PrintFunc()
	linkedList := NewLinkedList()

	linkedList.InsertToTail("newNode1")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode4")
	linkedList.InsertToTail("newNode5")
	linkedList.InsertToTail("newNode4")
	linkedList.InsertToTail("newNode3")
	linkedList.InsertToTail("newNode2")
	linkedList.InsertToTail("newNode1")

	fmt.Printf("result:%v\n", IsPalindrome(linkedList))
}
