package _1_graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	nodeCount int
	adj []*list.List
}
var found bool

func NewGraph(count int) *Graph {
	adj := make([]*list.List, count, count)

	for i:=0; i<count;i++  {
		adj[i] = &list.List{}
	}
	return &Graph{
		nodeCount:count,
		adj:adj,
	}
}
func (this *Graph)AddEdge(s int, t int)  {
	if s == t {
		return
	}
	this.adj[s].PushBack(t)
	this.adj[t].PushBack(s)
}

func (this *Graph)BFS(s int, t int)  {
	if s == t {return }
	visited := make([]bool, this.nodeCount)
	prev := make([]int, this.nodeCount)
	queue := list.New()

	for i:=1;i<this.nodeCount ;i++  {
		prev[i]=-1
	}
	visited[s] = true
	queue.PushBack(s)


	for queue.Len() > 0{
		qnode := queue.Front()
		queue.Remove(qnode)
		q := qnode.Value.(int)

		p := this.adj[q].Front()

		for  p!=nil {
			pval := p.Value.(int)
			if visited[pval] == true {
				p = p.Next()
				continue
			}
			prev[pval] = q
			visited[pval] = true
			if pval == t {
				this.printLine(prev, s, t)
				return
			}
			queue.PushBack(pval)
			p = p.Next()
		}
	}
}

func (this *Graph)DFS(s int, t int){
	if s == t {
		return
	}
	found = false
	visited := make([]bool, this.nodeCount)
	prev := make([]int, this.nodeCount)
	for i:=1;i<this.nodeCount ;i++  {
		prev[i]=-1
	}
	this.recurDfs(s, t, visited, prev)
	this.printLine(prev, s, t)

}

func (this *Graph)recurDfs(w int, t int, visited []bool, prev []int)  {
	if found == true {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	q := this.adj[w].Front()
	for  q!=nil {
		pval := q.Value.(int)
		if visited[pval] == true {
			q = q.Next()
			continue
		}
		prev[pval] = w

		this.recurDfs(q.Value.(int), t, visited, prev )
		q = q.Next()
	}
}
func (this *Graph)printLine(prev []int, s int, t int)  {
	if prev[t] != -1 && t != s{
		this.printLine(prev, s, prev[t])
	}
	fmt.Printf("%v ", t)
}

func (this *Graph)Print(){
	for i:=0;i<this.nodeCount ;i++  {
		fmt.Printf("%v:", i)
		p := this.adj[i].Front()
		for  p != nil {
			fmt.Printf("  %v ", p.Value)
			p = p.Next()
		}
		fmt.Printf("\n")
	}
}