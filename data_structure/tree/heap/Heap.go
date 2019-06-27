package heap

import (
	"container/list"
	"fmt"
)

type Heap struct {
	data []int //从下标1开始
	n int //可以存储的数量
	count int //现在存储的数量
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:make([]int, capacity + 1),
		n:capacity,
	}
}

func (heap *Heap)Insert(data int) bool {
	if heap.count >= heap.n {
		return false
	}
	heap.count++
	heap.data[heap.count] = data
	//自下往上堆化
	i := heap.count
	for i/2 > 0 && heap.data[i] > heap.data[i/2]{
		heap.data[i], heap.data[i/2] = heap.data[i/2], heap.data[i]
		i = i/2
	}
	return true
}

func (heap *Heap)RemoveMax() bool {
	if heap.count == 0 {
		return false
	}
	heap.data[1] = heap.data[heap.count]
	heap.count--
	heapify(heap.data, heap.count, 1)
	return true
}

//自上向下堆化
func heapify(data []int, n int, i int)  {
	for  {
		maxPos := i
		if 2 * i < n && data[2 * i] > data[i] {
			maxPos = 2 * i	
		}
		if (2 * i + 1) < n && data[2*i+1] > data[maxPos] {
			maxPos = 2 * i + 1
		}
		if maxPos == i {
			break
		}
		data[i], data[maxPos] = data[maxPos], data[i]
	}
}

func buildHeap(data []int, n int) {
	for i:= n/2;i>=1 ; i-- {
		heapify(data, n, i)
	}
}

func (heap *Heap)Sort(data []int){
	k := len(data) - 1
	buildHeap(data, k)
	for k > 1  {
		data[k], data[1] = data[1], data[k]
		heapify(data, k, 1)
		k--
	}
}

func (heap *Heap)Print()  {
	if heap.count < 1 {
		return
	}
	lists := list.New()
	lists.PushBack(1)
	len := 1
	i := 0
	for lists.Len() > 0 {
		postions := lists.Front()
		lists.Remove(postions)
		postion := postions.Value.(int)
		if 2 * postion <= heap.count {
			lists.PushBack(2 * postion)
		}
		if 2 * postion + 1 <= heap.count {
			lists.PushBack(2 * postion + 1)
		}
		fmt.Printf(" %v ", heap.data[postion])
		i++

		if len == i {
			i = 0
			len = lists.Len()
			fmt.Println()
		}
	}
}