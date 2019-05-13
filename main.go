package main

import (
	"github.com/xiaobudongzhang/sf/data_structure/tree/binary_tree"
)

func main()  {
	tree := binary_tree.NewTreeList()
	//tree.Root = binary_tree.NewNode(nil)
	tree.InitEle(&tree.Root)
	//tree.InOrderPrint()
	//var node *binary_tree.Node
	//node = tree.Root
	//tree.Root = binary_tree.NewNode("1")

	//fmt.Printf("this.root %v,%v\n", &tree.Root, tree.Root)
	//node.Left = binary_tree.NewNode("2")

	//fmt.Printf("node %v\n", node)
	//fmt.Printf("tree %v\n", tree.Root)
}
