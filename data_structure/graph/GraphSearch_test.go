package graph

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestGraphSearch_Bfs(t *testing.T) {
	lib.PrintFunc()
	graph := newGraphSearch(10)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	graph.Print()
}
