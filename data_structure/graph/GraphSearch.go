package graph

import (
	"container/list"
	"fmt"
)
var found bool
//无向图为基础
//todo 去重

type GraphSearch struct {
	size int
	adj []*list.List
}

func newGraphSearch(size int) *GraphSearch{
	adj := make([]*list.List, size+1)
	for k,_ :=range adj {
		adj[k] = &list.List{}
	}
	return &GraphSearch {
		size:size+1,
		adj:adj,
	}
}
func (graph *GraphSearch) AddEdge(s int, t int) bool {
	if s > graph.size || s < 0 || t > graph.size || t < 0 {
		return false
	}
	graph.adj[s].PushBack(t)
	graph.adj[t].PushBack(s)
	return true
}
//广度搜索
func (graph *GraphSearch) Bfs(s int, t int)  {
	if s == t {
		return
	}
	//定义辅助变量
	visited := make([]bool, graph.size + 1)//已经访问过的顶点
	queue := list.List{} //当前访问的顶点
	prev := make([]int , graph.size + 1)//改顶点从那个顶点来
	//初始化
	visited[s] = true
	for k,_ := range prev {
		prev[k] = -1
	}
	queue.PushBack(s)
	//查找
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		eleValue := ele.Value.(int)
		//遍历操作
		adjs := graph.adj[eleValue]
		current := adjs.Front()

		for current != nil {
			currentValue := current.Value.(int)

			if visited[currentValue] == true {//已经遍历过
				current = current.Next()
				continue
			}
			prev[currentValue] = eleValue
			if currentValue == t {
				graph.PrintSearch(prev, s, t)
				return
			}
			visited[eleValue] = true
			queue.PushBack(currentValue)

			current = current.Next()
		}
	}

}
//深度搜索
func (graph *GraphSearch) Dfs(s int, t int)  {
	if s == t {
		return
	}
	//定义辅助变量
	visited := make([]bool, graph.size + 1)//已经访问过的顶点
	prev := make([]int , graph.size + 1)//改顶点从那个顶点来
	found = false
	//初始化
	for k,_ := range prev {
		prev[k] = -1
	}
	graph.dfsRecur(s, t, visited, prev)

	graph.PrintSearch(prev, s, t)
}

func (graph *GraphSearch) dfsRecur(s int, t int, visited []bool, prev []int)  {
	if found == true {
		return
	}
	visited[s] = true

	if s == t {
		found = true
		return
	}

	//遍历操作
	adjs := graph.adj[s]
	current := adjs.Front()
	for current != nil {
		currentValue := current.Value.(int)

		if visited[currentValue] == true {//已经遍历过
			current = current.Next()
			continue
		}
		
		prev[currentValue] = s
		current = current.Next()
		graph.dfsRecur(currentValue, t, visited, prev)
	}

}

func (graph *GraphSearch) PrintSearch(prev []int, s int, t int) {

	if prev[t] != -1 && t != s  {
		graph.PrintSearch(prev, s, prev[t])
	}
	fmt.Printf("%v --> ", t)
}
func (graph *GraphSearch) Print() {
	for aEdge,ls :=range graph.adj {
		fmt.Printf("%v ==> ", aEdge)

		current := ls.Front()
		for current != nil {
			fmt.Printf(" %v ", current.Value.(int))
			current = current.Next()
		}

		fmt.Println()
	}
}
