package binary_search_tree

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestBinarySearchTree_Find(t *testing.T) {
	lib.PrintFunc()
	tree := NewBinaryNode()
	insertDatas := []int{13, 10, 9, 11, 16, 14}
	for _, data := range insertDatas {
		tree.Insert(data)
	}
	tree.InOrderPrint()
}

func TestBinarySearchTree_Delete(t *testing.T) {
	lib.PrintFunc()
	tree := NewBinaryNode()
	insertDatas := []int{13, 10, 9, 11, 16, 14}
	for _, data := range insertDatas {
		tree.Insert(data)
	}

	tree.InOrderPrint()
	fmt.Printf("\n------------14---------------\n")
	tree.Delete(14)
	tree.InOrderPrint()

	fmt.Printf("-------------16--------------\n")

	tree = NewBinaryNode()
	insertDatas = []int{13, 10, 9, 11, 16, 14}
	for _, data := range insertDatas {
		tree.Insert(data)
	}
	tree.Delete(16)
	tree.InOrderPrint()
	fmt.Printf("--------------11-------------\n")

	tree = NewBinaryNode()
	insertDatas = []int{13, 10, 9, 11, 16, 14}
	for _, data := range insertDatas {
		tree.Insert(data)
	}
	tree.Delete(11)
	tree.InOrderPrint()
	fmt.Printf("---------------------------\n")
}
