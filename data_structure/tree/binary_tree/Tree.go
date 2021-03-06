package binary_tree

import (
	"container/list"
	"fmt"
)

type Node struct {
	Ele   interface{}
	Left  *Node
	Right *Node
}

type TreeList struct {
	Root *Node
}

func NewTreeList() *TreeList {

	tree := &TreeList{
		nil,
	}
	return tree
}
func NewNode(ele interface{}) *Node {
	return &Node{
		Ele:   ele,
		Left:  nil,
		Right: nil,
	}
}

//递归前序创建
func (this *TreeList) InitEle(node **Node) {
	var inputEle string
	inputEle = "0"
	fmt.Scanf("%s", &inputEle)
	if inputEle == "100" {
		node = nil
		return
	}
	*node = NewNode(inputEle)
	fmt.Printf("  %s 的左节点\r\n", inputEle)
	this.InitEle(&((*node).Left))
	fmt.Printf(" %s 的右节点\r\n", inputEle)
	this.InitEle(&(*node).Right)
	return
}

//函数参数跟原指针是不同的指针，指向同一个地址，改变参数的指针指向不会改变原指针的指向
//所以传递指向指针的指针，这样的话，通过函数参数就能取到原指针了，从而进行修改
func (this *TreeList) InitEle2(node *Node) {
	var inputEle string
	inputEle = "0"
	fmt.Scanf("%s", &inputEle)
	if inputEle == "100" {
		node = nil
		return
	}
	node = NewNode(inputEle)
	fmt.Printf("  %s 的左节点\r\n", inputEle)
	this.InitEle2(node.Left)
	fmt.Printf("  %s 的右节点\r\n", inputEle)
	this.InitEle2(node.Right)
	return
}
func (this *TreeList) InOrderPrint() {
	lists := list.New()
	lists.PushBack(this.Root)
	i := 0
	qlen := lists.Len()

	for lists.Len() > 0 {
		ele := lists.Front()
		lists.Remove(ele)
		i++
		elev := ele.Value.(*Node)

		fmt.Printf("%v ", elev.Ele)
		if elev.Left != nil {
			lists.PushBack(elev.Left)
		}

		if elev.Right != nil {
			lists.PushBack(elev.Right)
		}

		if i == qlen { //这一层级走完
			i = 0
			qlen = lists.Len()
			fmt.Println("")
		}
	}
}

//self -> left -> right
func (this *TreeList) PreOrderPrint(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v", node.Ele)
	this.PreOrderPrint(node.Left)
	this.PreOrderPrint(node.Right)
}

func (this *TreeList) MidOrderPrint(node *Node) {
	if node == nil {
		return
	}
	this.MidOrderPrint(node.Left)
	fmt.Printf("%v", node.Ele)
	this.MidOrderPrint(node.Right)
}

func (this *TreeList) PostOrderPrint(node *Node) {
	if node == nil {
		return
	}
	this.PostOrderPrint(node.Left)
	this.PostOrderPrint(node.Right)
	fmt.Printf("%v", node.Ele)
}
