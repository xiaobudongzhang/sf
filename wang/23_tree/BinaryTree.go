package _3_tree

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root interface{}) *BinaryTree {
	return &BinaryTree{NewNode(root)}
}

