package _1_graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	nodeCount int
	adj       []*list.List
}

var found bool

func NewGraph(count int) *Graph {
	adj := make([]*list.List, count, count)

	for i := 0; i < count; i++ {
		adj[i] = &list.List{}
	}
	return &Graph{
		nodeCount: count,
		adj:       adj,
	}
}
func (this *Graph) AddEdge(s int, t int) {
	if s == t {
		return
	}
	this.adj[s].PushBack(t)
	//this.adj[t].PushBack(s)
}

func (this *Graph) BFS(s int, t int) {
	if s == t {
		return
	}
	visited := make([]bool, this.nodeCount)
	prev := make([]int, this.nodeCount)
	queue := list.New()

	for i := 1; i < this.nodeCount; i++ {
		prev[i] = -1
	}
	visited[s] = true
	queue.PushBack(s)

	for queue.Len() > 0 {
		qnode := queue.Front()
		queue.Remove(qnode)
		q := qnode.Value.(int)

		p := this.adj[q].Front()

		for p != nil {
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

func (this *Graph) DFS(s int, t int) {
	if s == t {
		return
	}
	found = false
	visited := make([]bool, this.nodeCount)
	prev := make([]int, this.nodeCount)
	for i := 1; i < this.nodeCount; i++ {
		prev[i] = -1
	}
	this.recurDfs(s, t, visited, prev)
	this.printLine(prev, s, t)

}

func (this *Graph) recurDfs(w int, t int, visited []bool, prev []int) {
	if found == true {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	q := this.adj[w].Front()
	for q != nil {
		pval := q.Value.(int)
		if visited[pval] == true {
			q = q.Next()
			continue
		}
		prev[pval] = w

		this.recurDfs(q.Value.(int), t, visited, prev)
		q = q.Next()
	}
}
func (this *Graph) printLine(prev []int, s int, t int) {
	if prev[t] != -1 && t != s {
		this.printLine(prev, s, prev[t])
	}
	fmt.Printf("%v ", t)
}

func (this *Graph) Print() {
	for i := 0; i < this.nodeCount; i++ {
		fmt.Printf("%v:", i)
		p := this.adj[i].Front()
		for p != nil {
			fmt.Printf("  %v ", p.Value)
			p = p.Next()
		}
		fmt.Printf("\n")
	}
}

func (this *Graph) TopoSort() {

	var inDegree []int
	inDegree = make([]int, this.nodeCount)
	for i := 0; i < this.nodeCount; i++ { //记录每个顶点的被依赖关系
		q := this.adj[i].Front()
		for q != nil {
			w := q.Value.(int)
			inDegree[w]++
			q = q.Next()
		}
	}

	queue := list.New()

	for i := 0; i < this.nodeCount; i++ { //没有被依赖关系的顶点
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)
		inode := node.Value.(int)
		fmt.Printf("->%v", inode)

		q := this.adj[inode].Front()
		for q != nil { //遍历这个节点对应的依赖减一
			w := q.Value.(int)
			inDegree[w]--
			if inDegree[w] == 0 {
				queue.PushBack(w)
			}
			q = q.Next()
		}
	}
}

func (this *Graph) TopoSortByDFS() {
	//先构建逆邻接表
	inverseAdj := make([]*list.List, this.nodeCount)
	for i := 0; i < this.nodeCount; i++ {
		inverseAdj[i] = list.New()
	}

	for i := 0; i < this.nodeCount; i++ {
		q := this.adj[i].Front()
		for q != nil { //遍历这个节点对应的依赖减一
			iq := q.Value.(int)
			inverseAdj[iq].PushBack(i)
			q = q.Next()
		}
	}

	visited := make([]bool, this.nodeCount)
	for i := 0; i < this.nodeCount; i++ { //深度优先遍历
		if visited[i] == false {
			visited[i] = true
			this.dfs(i, inverseAdj, visited)
		}
	}
}

func (this *Graph) dfs(i int, inverseAdj []*list.List, visited []bool) {

	q := this.adj[i].Front()
	for q != nil { //遍历这个节点对应的依赖减一
		iq := q.Value.(int)

		if visited[iq] == true {
			q = q.Next()
			continue
		}
		visited[iq] = true

		this.dfs(iq, inverseAdj, visited)

		q = q.Next()
	} //先把i这个顶点可达的所有顶点都打印出来之后，再打印它自己
	fmt.Printf("->%v", i)
}
