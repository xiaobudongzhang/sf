package _3_tree

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
	"fmt"
)

var compareFunc = func(v, node interface{}) int{
	vInt := v.(int)
	nodeInt := node.(int)

	if vInt > nodeInt{
		return 1
	}else if vInt < nodeInt{
		return -1
	}
	return 0
}

func TestBST_Find(t *testing.T) {
	lib.PrintFunc()
	bst := NewBST(10, compareFunc)
	bst.Insert(2)
	bst.Insert(11)
	bst.Insert(1)
	bst.Insert(3)

	findNode := bst.Find(2)
	fmt.Printf("findNode:%v", findNode.data)
}

func TestBST_Insert(t *testing.T) {
	lib.PrintFunc()
	bst := NewBST(10, compareFunc)
	bst.Insert(2)
	bst.Insert(11)
	bst.Insert(1)
	bst.Insert(3)

	bst.Print(bst.root)
}

func TestBST_Delete(t *testing.T) {
	lib.PrintFunc()
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(20)

	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(30)

	bst.Insert(2)
	bst.Insert(4)

	bst.Delete(1)
	fmt.Printf("-------------\n")
	bst.Print(bst.root)
	bst.Delete(3)
	fmt.Printf("-------------\n")
	bst.Print(bst.root)
	bst.Delete(20)
	fmt.Printf("-------------\n")
	bst.Print(bst.root)
}