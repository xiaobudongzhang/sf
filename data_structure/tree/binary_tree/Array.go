package binary_tree

import (
	"container/list"
	"fmt"
)

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
	//将当前节点加入队列，直到节点没有子树，打印出当前节点
	list := list.New()
	list.PushBack(1)//从根开始便利
	last := 1 //每一层的最后一个元素

	for list.Len() > 0 {
		ele := list.Front()
		list.Remove(ele)

		i := ele.Value.(int)
		fmt.Printf("%v ", this.nodes[i])
		if i == last {
			fmt.Println("")
		}
		tmplast := last
		if this.nodes[2 * i] != nil && 2 * i < this.nodesLen {
			if i == last {//下一层的最后一个元素一定是最后一个元素的子元素
				tmplast = 2 * i
			}
			list.PushBack(2 * i)
		}
		if this.nodes[2 * i + 1] != nil && (2 * i + 1) < this.nodesLen {
			if i == last {//下一层的最后一个元素一定是最后一个元素的子元素
				tmplast = 2 * i + 1
			}
			list.PushBack(2 * i + 1)
		}

		last = tmplast
	}
	fmt.Println("")
}

func (this Tree)Print(){
	fmt.Printf("array tree : %v \n", this)
}
