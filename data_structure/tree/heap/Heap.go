package heap

import (
	"container/list"
	"fmt"
)
type Comparator func(a, b interface{}) int

type Heap struct {
	data []interface{} //从下标1开始
	n int //可以存储的数量
	count int //现在存储的数量
	comparator Comparator
}

func NewHeap(capacity int, comparator Comparator) *Heap {
	return &Heap{
		data:make([]interface{}, capacity + 1),
		n:capacity,
		comparator:comparator,
	}
}

func (heap *Heap)Insert(data interface{}) bool {
	if heap.count >= heap.n {
		return false
	}
	heap.count++
	heap.data[heap.count] = data
	//自下往上堆化
	i := heap.count
	for i/2 > 0 && heap.comparator(heap.data[i], heap.data[i/2]) < 0{
		heap.data[i], heap.data[i/2] = heap.data[i/2], heap.data[i]
		i = i/2
	}
	return true
}
//由于不是二叉查找树，所以先直接遍历
func (heap *Heap)Update(oldData interface{},data interface{}) {
	for i:=0; i< heap.count; i++ {
		if heap.data[i] == oldData {
			heap.data[i] = data
		}
	}
	heap.heapify(heap.data, heap.count, 1)
}

func (heap *Heap) RemoveMin() interface{} {
	if heap.count == 0 {
		return nil
	}
	min := heap.data[1]


	heap.data[1] = heap.data[heap.count]
	heap.data[heap.count] = nil
	heap.count--
	heap.heapify(heap.data, heap.count, 1)



	return min
}
func (heap *Heap) Size() int {
	return heap.count
}

//自上向下堆化
func (heap *Heap) heapify(data []interface{}, n int, i int)  {
	for  {
		minPos := i
		if 2 * i <= n && heap.comparator(data[i], data[2 * i]) > 0{
			minPos = 2 * i
		}
		if (2 * i + 1) <= n && heap.comparator(data[minPos], data[2*i+1]) > 0 {
			minPos = 2 * i + 1
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
	}
}

func (heap *Heap)buildHeap(data []interface{}, n int) {
	for i:= n/2;i>=1 ; i-- {
		heap.heapify(data, n, i)
	}
}
//@todo 待验证
func (heap *Heap)Sort(data []interface{}){
	k := len(data) - 1
	heap.buildHeap(data, k)
	for k > 1  {
		data[k], data[1] = data[1], data[k]
		heap.heapify(data, k, 1)
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
func (heap *Heap) Data() []interface{}{
	return heap.data
}