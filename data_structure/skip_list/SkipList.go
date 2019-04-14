package skip_list

import (
	"math/rand"
	"fmt"
	"strconv"
)

//@todo ele不能重复
const MAX_LEVEL = 16
const P = 0.25

type SkipListNode struct {
	ele interface{}
	score float32
	backward *SkipListNode //在最0层查找前序节点使用
	level[] SkipListLevel //数组的
}

type SkipListLevel struct {
	forward *SkipListNode
	span int //跨度，快速计算出排名，从本节点到下一个节点的节点数
}

type SkipList struct {
		head *SkipListNode
		tail *SkipListNode
		level int
		length int
}

func (this *SkipList) Insert(ele interface{}, score float32) bool {
	update := [MAX_LEVEL]*SkipListNode{}//记录查找的路径，方便后面的操作
	rank := make([]int, MAX_LEVEL, MAX_LEVEL) //记录span跨度的


	cur := this.head
	//找到插入的元素
	//共level层，其中第0层是原始数据
	for i := this.level - 1; i >=0; i-- {

		if (i == this.level-1) {//第一级的为0
			rank[i] = 0
		} else {
			rank[i] = rank[i+1] //等于上一级
		}

		for cur.level[i].forward != nil && cur.level[i].forward.score < score {
			rank[i] += cur.level[i].span//加上每一步走的span
			cur = cur.level[i].forward
		}
		update[i] = cur
	}

	if cur.score == score {
		cur.ele = ele
		return true
	}
	lvl := RandomLevel()

	if lvl > this.level {//随机产生的层数大于现有的层数
		for i:= this.level;i <lvl ;i++  {
			rank[i] = 0
			update[i] = this.head  //指向下面的插入元素使用
			update[i].level[i].span = this.length //会被重新赋值,所以是什么值没影响
		}
		this.level = lvl
	}
	//fmt.Printf("lvl:%v-- update:%v\n", lvl,update)
	newNode := newNode(lvl, score, ele)
	for i:=0; i < lvl; i++  {
		newNode.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = newNode

		/**
		 * 总的减去前面的，如果是第0层则为1，因为update[0].level[0] = 1
		 */
		newNode.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	/* increment span for untouched levels */
	for i := lvl; i < this.level; i++ {
		update[i].level[i].span++
	}
	//更新新节点的前序节点
	if update[0] == this.head {
		newNode.backward = nil
	} else {
		newNode.backward = update[0]
	}

	if newNode.level[0].forward == nil {
		this.tail = newNode
	} else {
		newNode.level[0].forward.backward = newNode //更新下一个节点的前序节点
	}

	this.length++
	return true
}
func (this *SkipList) Delete(ele interface{}, score float32) bool {
	update := [MAX_LEVEL]*SkipListNode{}//记录查找的路径，方便后面的操作
	cur := this.head
	//找到删除的元素
	//共level层，其中第0层是原始数据
	//找到前驱节点
	for i := this.level - 1; i >=0; i-- {
		for cur.level[i].forward != nil && (cur.level[i].forward.score < score) {
			cur = cur.level[i].forward
		}
		update[i] = cur
	}
	findCur := cur.level[0].forward
	if findCur != nil && findCur.ele != ele {//没找到
		return true
	}
	//找到后删除
	for i:=0;i< this.level;i++{
		if (update[i].level[i].forward == findCur) {//如果是跟删除的元素相连的话
			update[i].level[i].span += findCur.level[i].span - 1 //span 进行扩充
			update[i].level[i].forward = findCur.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}

	if findCur.level[0].forward != nil {//backward修改
		findCur.level[0].forward.backward = findCur.backward //当前的下一个指向下当前的前一个
	} else {
		this.tail = findCur.backward
	}
	//对于没有索引的层，减一，直到为1(根据head的forward判断)
	for this.level > 1 && this.head.level[this.level-1].forward == nil {
		this.level--
	}

	this.length--

	return true
}
func (this *SkipList) Get(ele interface{})  interface{}{
	return nil
}

func (this *SkipList) GetRank(ele interface{})  int {
	return 0
}

func  RandomLevel() int {
	level := 1
	//保证了每隔1/p次产生1个level,每隔（1/p）的平方产生2个level........
	//为什么能保证，因为连续产生n个level的概率是p的n次方，所以1/p的n次方产生一个连续n次的level
	//看似简单，理解起来还是挺复杂的，设计的很巧妙，通过随机函数保证了树的形状
	for ;float32(rand.Intn(0xFFFF)) < P * 0xFFFF ;  {
		level++
	}
	if level > MAX_LEVEL {
		return MAX_LEVEL
	}
	return level
}
func (this *SkipList) Print()  {
	head := this.head
	for i:=this.level -1; i>=0; i--  {

		//头节点
		fmt.Printf("|")
		for j:=0; j < head.level[i].span; j++ {
			fmt.Printf(" ")
		}
		//计算前面数字占用的长度
		if head.level[i].forward != nil {
			nextScore := head.level[i].forward.score
			cur1 := head
			for cur1.score < nextScore {
				for x:=0; x< len(strconv.Itoa(int(cur1.score)));x++  {
					fmt.Printf(" ")
				}
				cur1 = cur1.level[0].forward
			}
		}



		cur := head.level[i].forward
		for cur != nil  {
			fmt.Printf("%v",cur.score)

			for j:=0; j< cur.level[i].span; j++ {
				fmt.Printf(" ")
			}

			//计算前面数字占用的长度,从上一个到下一个中间的
			if cur.level[i].forward != nil {
				cur11 := cur.level[0].forward
				nextScore2 := cur.level[i].forward.score

				for cur11.score > cur.score && cur11.score < nextScore2 {
					for x:=0; x< len(strconv.Itoa(int(cur11.score)));x++  {
						fmt.Printf(" ")
					}
					cur11 = cur11.level[0].forward
				}
			}


			cur = cur.level[i].forward
		}
		fmt.Printf("\n")
	}
}

func  newNode(level int, score float32, ele interface{}) *SkipListNode {
	levels := make([]SkipListLevel, level,level)
	return &SkipListNode{
		ele:ele,
		score:score,
		level:levels,
	}
}

func NewSkipList() *SkipList{
	head := newNode(MAX_LEVEL, 0, nil)

	for i:=0;i<MAX_LEVEL ;i++  {
		head.level[i].span = 0
		head.level[i].forward = nil
	}
	return &SkipList{
		head:head,
		tail:nil,
		level:1,
		length:0,
	}
}

