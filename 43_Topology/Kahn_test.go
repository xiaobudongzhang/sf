package _3_Topology

import (
	"31_graph"
	"fmt"
	"lib"
	"testing"
)

func TestGraph_TopoSort(t *testing.T) {
	lib.PrintFunc()

	graph := _1_graph.NewGraph(10)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)

	graph.AddEdge(2, 5)

	graph.AddEdge(3, 4)

	graph.AddEdge(4, 5)

	graph.AddEdge(5, 7)

	graph.AddEdge(6, 4)
	graph.AddEdge(6, 7)

	graph.AddEdge(9, 8)

	graph.Print()

	graph.TopoSort()

	fmt.Printf("\n---------------\n")

	graph.TopoSortByDFS()
}
