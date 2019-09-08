package graph

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/data_structure/graph"
)

type Dfs struct {
	graph *graph.GraphList
	size int
}

func NewDfs(size int) *Dfs {
	return &Dfs{
		graph:graph.NewGraphList(size),
		size:size+1,
	}
}

func (this *Dfs) TopoSort()  {
	//构建逆邻接表
	inverseAdj := graph.NewGraphList(this.size)
	graph := this.graph
	for  kp,v := range graph.Data() {
		current := v
		for current != nil {
			k := current.Data()
			if k == -1 {
				current = current.Next()
				continue
			}

			inverseAdj.AddEdge(k, kp)
			current = current.Next()
		}
	}

	//递归打印改顶点所依赖的所有顶点后再打印自己
	visited := make([]bool, this.size + 1)
	for k,_ :=range inverseAdj.Data() {
		if visited[k] {
			continue
		}
		visited[k] = true
		this.dfs(k, inverseAdj, visited)
	}
}

func (this *Dfs) dfs(v int, inverseAdj *graph.GraphList, visited []bool) {
	current := inverseAdj.Data()[v]
	//如果这个顶点有依赖才打印

	for current != nil {
		k := current.Data()
		if k == -1 {
			current = current.Next()
			continue
		}


		if visited[k] == true {
			current = current.Next()
			continue
		}
		visited[k] = true
		this.dfs(k, inverseAdj, visited)
		current = current.Next()
	}
	//先把i这个顶点可达的所有顶点都打印出来之后，再打印它自己
	if inverseAdj.InList(v) == true {
		fmt.Printf("%v --> ", v)
	}
}
