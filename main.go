package main

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/data_structure/tree/binary_tree"
)

func main()  {
	tree := binary_tree.NewTreeList()

	tree.InitEle(&tree.Root)
	fmt.Println("----------inorder-------------------")
	tree.InOrderPrint()
	fmt.Println("-------------------------------------")

	fmt.Println("----------preorder-------------------")
	tree.PreOrderPrint(tree.Root)
	fmt.Println("")
	fmt.Println("-------------------------------------")

	fmt.Println("----------midorder-------------------")
	tree.MidOrderPrint(tree.Root)
	fmt.Println("")
	fmt.Println("-------------------------------------")

	fmt.Println("----------postorder-------------------")
	tree.PostOrderPrint(tree.Root)
	fmt.Println("")
	fmt.Println("-------------------------------------")
}
