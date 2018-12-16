package _7_skiplist

import (
	"math/rand"
	"fmt"
)

const MAX_LEVEL = 15;

type Node struct {
	val int
	level int
	forward []*Node//数组的下标表示第几层，数组的值是指针指向下一级，设计的很巧妙
}

type SkipList struct {
	head *Node
	level int
	length int
}

func newNode(v int,level int) *Node {
	return &Node{val:v,level:level,forward:make([]*Node, level,level)}
}

func NewSkipList() *SkipList{
	return &SkipList{
		head:newNode(0, MAX_LEVEL),
		level: 0,
		length: 0,
		}
}

func (this *SkipList) Length() int {
	return this.length
}

func (this *SkipList) Level() int{
	return this.level
}

func (this *SkipList) Insert(v int)  {
	cur := this.head
	update := [MAX_LEVEL]*Node{}
	//存放插入节点的前面节点
	for i:= 0 ;i < MAX_LEVEL; i++ {
		update[i] = cur
	}

	//保存老的路径
	level := randomLevel()
	fmt.Printf("%v", level)
	for i := level -1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < v{
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	//插入节点
	newNode := newNode(v, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	if level > this.level{
		this.level = level
		this.length = 1
	} else if this.level == level{
		this.length++
	}
}

func randomLevel() int {
	level := 1
	for i:=1;i<MAX_LEVEL ;i++  {
		rand := rand.Int()
		if rand % 2 == 1{
			level++
		}
	}
	return level
}

func (this *SkipList) Find(v int) *Node {
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < v{
			cur = cur.forward[i]
		}
	}

	if cur.forward[0] != nil && cur.forward[0].val == v {
		return cur.forward[0]
	}
	return nil
}

func (this *SkipList) Delete(v int) {
	//找到要删除数据的节点与前一级节点
	prev := [MAX_LEVEL]*Node{}
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil &&  cur.forward[i].val < v{
				cur = cur.forward[i]
			}
			fmt.Printf("cur:%v--", cur.val)
			prev[i] = cur
	}
	if cur.forward[0].level == this.level{
		this.length--
	}
	//根据删除节点的信息删除
	for i := this.level -1; i >= 0; i-- {
		if prev[i].forward[i] != nil && prev[i].forward[i].val == v{
			prev[i].forward[i] = prev[i].forward[i].forward[i]
		}
	}
	//this.length--
	if this.length == 0{
		//重新计算length
		cur := this.head
		for i:= this.level -2;i>=0 ;i--  {
			for cur.forward[i] != nil{
				this.length++
				cur = cur.forward[i]
			}

			if this.length >0 {//当前级有数据
				this.level = i+1
				break
			}else {//当前级也没数据了，等于下一级先
				this.level = i
			}
		}
	}
}


func (this *SkipList) Print()  {
	fmt.Printf("\n---level:%v:height:%v\n",this.level, this.length)
	for i := this.level - 1; i >= 0; i-- {
		fmt.Printf("l:%v ",i)
		cur := this.head
		for cur.forward[i] != nil {
			fmt.Printf(" %v",cur.forward[i].val)
			cur = cur.forward[i]
		}
		fmt.Printf("---\n:")
	}
}
