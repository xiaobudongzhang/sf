package _8_heap

import (
	"09_queue"
	"fmt"
)

type Heap struct {
	data     []int
	maxCount int
	count    int
}

func NewHeap(maxCount int) *Heap {
	return &Heap{
		maxCount: maxCount,
		count:    0,
		data:     make([]int, maxCount),
	}
}

func (this *Heap) Insert(val int) {
	if this.count >= this.maxCount {
		return
	}
	this.count++
	this.data[this.count] = val
	//堆化(从下往上)
	i := this.count
	for (i/2 > 0) && this.data[i] > this.data[i/2] {
		this.data[i], this.data[i/2] = this.data[i/2], this.data[i]
		i = i / 2
	}

}

func (this *Heap) RemoveMax() {
	this.data[1] = this.data[this.count]
	this.count--
	this.Heap(this.count, 1)
}

//从上往下
func (this *Heap) Heap(count int, pos int) {
	p := pos
	//找到最大位置（自己与两个子类），交换
	for {
		maxPos := p
		if 2*p <= count && this.data[2*p] > this.data[maxPos] {
			maxPos = 2 * p
		}

		if 2*p+1 <= count && this.data[2*p+1] > this.data[maxPos] {
			maxPos = 2*p + 1
		}

		//如果最大节点是自己，说明交换结束
		if maxPos == p {
			break
		}
		this.data[maxPos], this.data[p] = this.data[p], this.data[maxPos]
		p = maxPos
	}
}
func (this *Heap) Print() {
	q := _9_queue.NewLinkedListQueue()

	p := 1
	q.EnQueue(p)
	qlen := q.Len()
	i := 0
	height := 0
	fmt.Printf("level:0--")
	for !q.IsEmpty() {
		p := q.DeQueue().(int)
		fmt.Printf("%+v", this.data[p])
		if p/2 > 0 {
			fmt.Printf("(%v) ", this.data[p/2])
		}
		if (2 * p) <= this.count {
			q.EnQueue(2 * p)
		}

		if ((2 * p) + 1) <= this.count {
			q.EnQueue((2 * p) + 1)
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
}
