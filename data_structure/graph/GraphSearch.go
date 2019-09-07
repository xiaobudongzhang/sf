package graph

import (
	"container/list"
	"fmt"
)

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
func (graph *GraphSearch) Bfs(s int, t int) bool {
	return true
}
//深度搜索
func (graph *GraphSearch) Dfs(s int, t int) bool {
	return true
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
