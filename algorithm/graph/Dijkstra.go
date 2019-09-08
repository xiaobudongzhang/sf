package graph

import (
	"container/list"
	"fmt"
	"github.com/xiaobudongzhang/sf/data_structure/tree/heap"
	"math"
)
type Graph struct {
	size int
	adj       []*list.List
}
type Edge struct {
	sid int
	tid int
	w   int //权重
}

type Vertex struct {
	id   int
	dist int //从起始顶点到这个顶点的距离
}

var found bool

func NewGraph(size int) *Graph {
	adj := make([]*list.List, size+1)
	for k,_ :=range adj {
		adj[k] = list.New()
	}
	return &Graph{
		size: size+1,
		adj:adj,
	}
}

func NewEdge(sid int, tid int, w int) *Edge {
	return &Edge{
		sid:sid,
		tid:tid,
		w:w,
	}
}
func newVertex(id int, dist int) *Vertex {
	return &Vertex{
		id:id,
		dist:dist,
	}
}

func (this *Graph) AddEdge(s int, t int, w int) {
	this.adj[s].PushBack(NewEdge(s, t, w))
}

func (this *Graph) Dijkstra(s int, t int) {
	//初始化

	inverseIntComparator := func(a, b interface{}) int {
		if a.(*Vertex).dist > b.(*Vertex).dist {
			return 1
		} else {
			return -1
		}
	}

	queue := heap.NewHeap(this.size+1, inverseIntComparator)
	inqueue := make([]bool, this.size+1)
	prev := make([]int, this.size+1)
	vertexs := make([]*Vertex, this.size+1) //起始顶点到每个顶点的距离

	//初始化为无穷大
	for k,_ := range vertexs {
		vertexs[k] = newVertex(k, math.MaxInt16)
	}
	vertexs[s].dist = 0;
	inqueue[s] = true
	queue.Insert(vertexs[s])

	//查找最短路径
	for queue.Size() > 0{
		minVertex := queue.RemoveMin().(*Vertex)

		if minVertex.id == t  {
			break
		}
		current := this.adj[minVertex.id].Front()

		for current != nil {//遍历相邻的顶点
			e := current.Value.(*Edge)
			//取出一条minVertex相连的边
			nextVertex := vertexs[e.tid]

			//找一条到nextVertex更短的路径,多个点都到一个点时
			if minVertex.dist + e.w < nextVertex.dist  {
				oldNextVertex := nextVertex
				//因为这里是引用所以队列里的值已经改变了
				//@todo 传递值
				nextVertex.dist = minVertex.dist + e.w
				prev[nextVertex.id] = minVertex.id
				//如果queue里的值更小更新queue里的值
				if inqueue[nextVertex.id] ==  true {
					queue.Update(oldNextVertex, nextVertex)
				} else {
					queue.Insert(nextVertex)
					inqueue[nextVertex.id] = true
				}
			}
			current = current.Next()
		}
	}

	fmt.Printf("%v -> ", s)
	this.print(s,t,prev)
}

func (this *Graph) print(s int, t int, prev []int) {
	if s==t {
		return
	}
	this.print(s,prev[t],prev)
	fmt.Printf("%v -> ", t)
}
