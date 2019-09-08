package graph

import (
	"container/list"
	"fmt"
	"github.com/xiaobudongzhang/sf/data_structure/graph"
)

//拓扑排序
type Kahn struct {
	graph *graph.GraphList
	size int
}

func NewKahn(size int) *Kahn {
	graph := graph.NewGraphList(size)
	return &Kahn{
		graph:graph,
		size:size+1,
	}
}
func (this *Kahn) TopoSort()  {
	graph := this.graph
	//统计每个顶点的入度
	inDegree := make([]int, this.size)
	for k,_ :=range inDegree {
		inDegree[k] = -1
	}
	for ki, v := range graph.Data() {
		current := v
		for  current != nil {
			k := current.Data()
			if k == -1 {
				current = current.Next()
				continue
			}
			if inDegree[ki] == -1 {
				inDegree[ki] = 0
			}
			if inDegree[k] == -1 {
				inDegree[k] = 0
			}
			inDegree[k]++
			current = current.Next()
		}
	}

	//将入度为0的顶点，加入到输出队列中
	queue := list.List{}
	for k,v :=range inDegree {
		if v == 0 {
			fmt.Printf("v:%v\n", k )
			queue.PushBack(k)
		}
	}

	//输出
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		eleV := ele.Value.(int)
		fmt.Printf("%v --> ", eleV)
		//将相关的依赖顶点入度减1
		current := graph.Data()[eleV]
		for current != nil {
			k := current.Data()
			if k == -1 {
				current = current.Next()
				continue
			}
			inDegree[k]--
			if inDegree[k] == 0 {
				queue.PushBack(k)
			}
			current = current.Next()
		}
	}
}

