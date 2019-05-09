package binary_tree

import "fmt"

type Tree struct {
	size int
	nodesLen int
	nodes[]interface{}
}

func NewTree() *Tree {
	//@todo 先默认1000，待自动拓展
	nodes := make([]interface{}, 1000, 1000)
	return &Tree{
		size:1,
		nodes:nodes,
		nodesLen:1000,
	}
}

//第一个元素从1开始
func (this *Tree)Insert(data interface{}) bool {

	if this.size > this.nodesLen - 1 {
		return false
	}
	this.nodes[this.size] = data
	this.size++
	return true
}
//按层遍历
func (this *Tree)InOrderPrint()  {

}

func (this Tree)Print(){
	fmt.Printf("array tree : %v \n", this)
}
