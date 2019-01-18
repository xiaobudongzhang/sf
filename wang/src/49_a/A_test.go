package _9_a

import (
	"testing"
	"lib"
)

func TestGraph_Astar(t *testing.T) {
	lib.PrintFunc()

	graph := NewGraph(14)

	graph.AddVetex(0,320,630)
	graph.AddVetex(1,300,630)
	graph.AddVetex(2,280,625)
	graph.AddVetex(3,270,630)
	graph.AddVetex(4,320,700)
	graph.AddVetex(5,360,620)
	graph.AddVetex(6,320,590)
	graph.AddVetex(7,370,580)
	graph.AddVetex(8,350,730)
	graph.AddVetex(9,390,690)
	graph.AddVetex(10,400,620)
	graph.AddVetex(11,420,580)
	graph.AddVetex(12,270,670)
	graph.AddVetex(13,270,600)





	graph.AddEdge(0, 4, 60)
	graph.AddEdge(0, 6, 60)
	graph.AddEdge(0, 5, 60)
	graph.AddEdge(0, 1, 20)
	graph.AddEdge(1, 2, 20)
	graph.AddEdge(2, 3, 10)

	graph.AddEdge(3, 12, 40)
	graph.AddEdge(3, 13, 30)

	graph.AddEdge(4, 8, 50)
	graph.AddEdge(6, 7, 70)

	graph.AddEdge(5, 8, 70)
	graph.AddEdge(5, 10, 50)
	graph.AddEdge(5, 9, 80)
	graph.AddEdge(7, 11, 50)
	graph.AddEdge(8, 9, 50)

	graph.AddEdge(9, 10, 60)
	graph.AddEdge(11, 10, 60)
	graph.AddEdge(12, 4, 40)
	graph.AddEdge(13, 6, 50)


	graph.Astar(0,10)
}
