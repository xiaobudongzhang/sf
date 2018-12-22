package _3_tree

import (
	"09_queue"
	"fmt"
	"math"
)

//计算树的高度....
func (this *BST) Height(p *Node) int{
	lHeight := 0
	rHeight := 0

	if p == nil {
		return  0
	}

	lHeight = this.Height(p.left)
	rHeight = this.Height(p.right)

	if lHeight > rHeight {
		return lHeight + 1
	}else {
		return rHeight + 1
	}
}

func (this *BST) Height2(p *Node) float64{
	if p == nil {
		return 0
	}
	return math.Max(this.Height2(p.left), this.Height2(p.right)) + 1
}

//按层遍历(返回树的高度)

func (this *BST)InOrder() int{
	p := this.root
	q := _9_queue.NewLinkedListQueue()

	q.EnQueue(p)

	qlen := q.Len()
	i := 0
	height := 0
	fmt.Printf("level:0--")
	for !q.IsEmpty(){
		p := q.DeQueue().(*Node)
		fmt.Printf("%+v ", p.data)

		if p.left != nil {
			q.EnQueue(p.left)
		}

		if p.right != nil {
			q.EnQueue(p.right)
		}

		i++
		if i == qlen {
			height++
			fmt.Printf("\n")
			qlen = q.Len()
			if qlen > 0 {
				fmt.Printf("level:%v--", i)
			}
			i = 0
		}
	}
	fmt.Printf("height:%v\n", height)
	return height
}

