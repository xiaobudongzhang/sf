package _3_tree

import (
	"fmt"
	"lib"
	"testing"
)

func TestBinaryTree_InOrder(t *testing.T) {
	lib.PrintFunc()
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(20)

	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(30)

	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(50)
	bst.Insert(25)
	bst.InOrder()

	height := bst.Height(bst.root)
	fmt.Printf("height:%v\n", height)

	height2 := bst.Height2(bst.root)
	fmt.Printf("height2:%v\n", height2)
}
