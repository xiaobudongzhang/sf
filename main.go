package main

import "github.com/xiaobudongzhang/sf/data_structure/tree/binary_tree"

func main()  {
	tree := binary_tree.NewTreeList()
	tree.InitEle(tree.Root)
	tree.InOrderPrint()
}
