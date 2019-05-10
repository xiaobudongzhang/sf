package binary_tree

import "testing"

func TestTree_Print(t *testing.T) {
	tree := NewTree()
	tree.Insert("A")
	tree.Insert("B")
	tree.Insert("C")
	tree.Insert("D")
	tree.Insert("E")
	tree.Insert("F")
	tree.Insert("G")
	tree.Insert("H")
	tree.Insert("I")
	tree.Insert("J")

	tree.Print()


	tree.InOrderPrint()
}
