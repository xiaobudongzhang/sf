package graph

import "fmt"

//邻接表
//@todo 检查是否重复
type node struct {
	data int
	next *node
}

type GraphList struct {
	datas []*node
	size int
}

func newNode(data int) *node {
	return &node{data:data}
}

func NewGraphList (size int) *GraphList {
	datas := make([]*node, size + 1)
	for k, _ :=range datas {
		tmpNode := newNode(-1)//初始化时重新newNode
		datas[k] = tmpNode
	}
	return &GraphList {
		datas:datas,
		size:size,
	}
}

func (graph *GraphList) AddEdge(a int, b int) bool {
	if a > graph.size || a < 1 || b < 1 {
		return false
	}
	graph.addNode(a, b)
	return true
}

func (graph *GraphList) RemoveEdge(a int , b int) bool {
	if a > graph.size || a < 1 || b < 1 {
		return false
	}
	graph.removeNode(a, b)
	return true
}

func (graph *GraphList) addNode(edge int, addData int) bool {
	node := newNode(addData)
	current := graph.datas[edge]

	for current.next != nil {
		current = current.next
	}

	current.next = node
	return true
}

func (graph *GraphList) removeNode(edge int, removeData int) bool {
	removePrev := graph.datas[edge]
	find := false
	for removePrev != nil {
		if removePrev.next != nil && removePrev.next.data == removeData {
			find = true
			break
		}
		//逻辑写在上面
		removePrev = removePrev.next
	}
	if find == false {
		return true
	}
	if removePrev.next != nil {
		removePrev.next = removePrev.next.next
		return true
	}
	removePrev.next = nil
	return true
}

func (graph *GraphList) Print() {
	for aEdge, data :=range graph.datas {
		fmt.Printf("aEdge:%v => ", aEdge)

		current := data
		for current != nil {
			fmt.Printf(" %v ", current.data)

			current = current.next
		}
		fmt.Println()
	}
}