package _9_a

import (
	"container/list"
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	"math"
)

type Graph struct {
	nodeCount int
	adj       []*list.List
	vertexs   []*Vertex
}
type Edge struct {
	sid int
	tid int
	w   int //权重
}

type Vertex struct {
	id   int
	dist int
	f    int
	x    int
	y    int
}

var found bool

func NewGraph(count int) *Graph {
	adj := make([]*list.List, count, count)
	vertexs := make([]*Vertex, count)
	for i := 0; i < count; i++ {
		adj[i] = &list.List{}
		vertexs[i] = nil
	}
	return &Graph{
		nodeCount: count,
		adj:       adj,
		vertexs:   vertexs,
	}
}

func NewEdge(sid int, tid int, w int) *Edge {
	return &Edge{
		sid: sid,
		tid: tid,
		w:   w,
	}
}
func NewVertex(id int, x int, y int) *Vertex {
	return &Vertex{
		id:   id,
		x:    x,
		y:    y,
		f:    math.MaxInt32,
		dist: math.MaxInt32,
	}
}

func (this *Graph) AddEdge(s int, t int, w int) {

	this.adj[s].PushBack(*NewEdge(s, t, w))
}

func (this *Graph) AddVetex(id int, x int, y int) {
	this.vertexs[id] = NewVertex(id, x, y)
}

func (this *Graph) Astar(s int, t int) {
	var predecessor []int
	predecessor = make([]int, this.nodeCount) //用来还原最短路径
	vertexs := this.vertexs                   //起始顶点到这个顶点的距离

	//小顶堆
	// Max-heap
	queue := binaryheap.NewWithIntComparator() // empty (min-heap)
	inverseIntComparator := func(a, b interface{}) int {
		if a.(Vertex).f > b.(Vertex).f {
			return 1
		} else {
			return -1
		}
	}
	queue = binaryheap.NewWith(inverseIntComparator) // empty (min-heap)

	inQueue := make([]bool, this.nodeCount) //标记是否进入过队列
	vertexs[s].dist = 0
	vertexs[s].f = 0
	inQueue[s] = true

	queue.Push(*vertexs[s]) //先把起始顶点放到队列中

	for !queue.Empty() {
		minVertex, _ := queue.Pop()
		//fmt.Printf("minVertex:%v\n", minVertex)
		//break
		if minVertex.(Vertex).id == t {
			break
		} //找到了最短路径

		e := this.adj[minVertex.(Vertex).id].Front()

		//fmt.Printf("e:%v\n",e)
		for e != nil {
			ep := e.Value.(Edge)
			//取出一条minVertex相连的边
			nextVertex := vertexs[ep.tid]

			//找一条到nextVertex更短的路径,多个点都到一个点时
			if minVertex.(Vertex).dist+ep.w < nextVertex.dist {
				nextVertex.dist = minVertex.(Vertex).dist + ep.w //更新dist
				nextVertex.f = nextVertex.dist + hManhattan(*nextVertex, *vertexs[t])

				predecessor[nextVertex.id] = minVertex.(Vertex).id
				//如果queue里的值更小更新queue里的值
				queue.Push(*nextVertex)
				inQueue[nextVertex.id] = true

			}

			if nextVertex.id == t {
				break
			}
			e = e.Next()
		}

		fmt.Printf("queue:%v\n", queue.Values())

	}
	this.print(s, t, predecessor)
}

func (this *Graph) print(s int, t int, processor []int) {
	if s == t {
		return
	}
	this.print(s, processor[t], processor)
	fmt.Printf("->%v", t)
}

func hManhattan(v1 Vertex, v2 Vertex) int {
	return int(math.Abs(float64(v1.x)-float64(v2.x))) + int(math.Abs(float64(v1.y)-float64(v2.y)))
}
