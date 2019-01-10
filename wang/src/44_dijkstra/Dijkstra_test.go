package _4_dijkstra

import (
	"testing"
	"lib"
)

func TestGraph_Dijkstra(t *testing.T) {
	lib.PrintFunc()

	graph := NewGraph(6)

	graph.AddEdge(0,1,10)
	graph.AddEdge(0,4,15)

	graph.AddEdge(1,2,15)
	graph.AddEdge(1,3,2)

	graph.AddEdge(2,5,5)

	graph.AddEdge(3,2,1)
	graph.AddEdge(3,5,12)

	graph.AddEdge(4,5,10)



	graph.Dijkstra(0,5)
}
