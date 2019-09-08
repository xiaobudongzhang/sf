package graph

import "fmt"

//邻接表
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
	//检查是否已经添加
	if graph.exitNode(a, b) == true {
		return true
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
func (graph *GraphList) exitNode(edge int, addData int) bool {
	current := graph.datas[edge]
	for current != nil {
		if current.data == -1 {
			current = current.next
			continue
		}
		if current.data == addData {
			return true
		}

		current = current.next
	}
	return false
}

func (graph *GraphList) InList(edge int) bool {

	//有子类的顶点
	current := graph.datas[edge]
	for current != nil {
		if current.data == -1 {
			current = current.next
			continue
		}
		return true
		current = current.next
	}
	//改顶点在子类中
	for _,vp :=range graph.datas  {
		current := vp
		for current != nil {
			if current.data == -1 {
				current = current.next
				continue
			}
			if current.data == edge {
				return true
			}
			current = current.next
		}
	}

	return false
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

func (graph *GraphList) Data() []*node {
	return graph.datas
}


func (node *node) Next() *node {
	return node.next
}

func (node *node) Data() int {
	return node.data
}