package _3_tree

import "fmt"

type Tree struct {
	size int
	nodes[100]interface{}
}
func NewTree(datas[]interface{}) *Tree {
	tree := &Tree{
		size:len(datas),
	}
	for i := 0; i < len(datas); i++ {
		tree.nodes[i] = datas[i]
	}
	return tree
}

func (this *Tree)Print(index int)  {
	if index < 0 || index > this.size{
		return
	}
	//left child
	this.Print(index << 1)
	fmt.Printf("[%d]:%v\n", index, this.nodes[index])
	//right
	this.Print((index << 1) + 1)
}