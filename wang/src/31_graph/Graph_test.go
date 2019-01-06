package _1_graph

import (
	"testing"
	"lib"
	"fmt"
)

func TestGraph_AddEdge(t *testing.T) {
	lib.PrintFunc()
	graph := NewGraph(10)

	graph.AddEdge(0,1)
	graph.AddEdge(0,3)

	graph.AddEdge(1,2)
	graph.AddEdge(1,4)

	graph.AddEdge(2,5)

	graph.AddEdge(3,4)

	graph.AddEdge(4,5)

	graph.AddEdge(5,7)

	graph.AddEdge(6,4)
	graph.AddEdge(6,7)

	graph.Print()

	graph.BFS(0,6)
	fmt.Printf("\n-------------\n")
	graph.DFS(0,6)
}

func TestGraph_BFS(t *testing.T) {
	lib.PrintFunc()

}

func TestGraph_TopoSort(t *testing.T) {
	lib.PrintFunc()


}