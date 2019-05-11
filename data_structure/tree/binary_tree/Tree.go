package binary_tree

import (
	"container/list"
	"fmt"
)

var inputEle string

type Node struct {
	ele interface{}
	left *Node
	right *Node
}

type TreeList struct {
	Root *Node
	Size int
}
func NewTreeList() *TreeList {

	tree := &TreeList{
		nil,
		0,
	}
	return tree
}
func NewNode(ele interface{}) *Node {
	return &Node{
		ele:ele,
	}
}


func (this *TreeList)InitEle(node *Node) {
	fmt.Scanf("%s", &inputEle)
	if inputEle == "end" {
		node = nil
	} else {
		newNode := NewNode(inputEle)
		if newNode == nil {
			return
		}
		fmt.Printf(" created %s 的左节点\r\n", inputEle)
		this.InitEle(newNode.left)
		fmt.Printf("created %s 的右节点\r\n", inputEle)
		this.InitEle(newNode.right)
	}
	return
}

func (this *TreeList)InOrderPrint()  {
	lists := list.New()
	lists.PushBack(this.Root)
	i := 0
	qlen := lists.Len()

	for lists.Len() > 0 {
		ele := lists.Front()
		lists.Remove(ele)

		i++
		elev := ele.Value.(*Node)

		fmt.Printf("%v", elev.ele)
		if elev.left != nil {
			lists.PushBack(elev.left)
		}

		if elev.right != nil {
			lists.PushBack(elev.right)
		}

		if i == qlen {//这一层级走完
			i = 0
			qlen = lists.Len()
			fmt.Println("")
		}
	}
}
