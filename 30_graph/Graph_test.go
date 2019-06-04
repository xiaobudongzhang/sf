package _0_graph

import (
	"lib"
	"testing"
	//"fmt"
)

func TestGraph_InsertA(t *testing.T) {
	lib.PrintFunc()
	graph := NewGraph(10)

	graph.InsertA(1)
	graph.InsertA(2)
	graph.InsertA(3)

	graph.InsertOfA(1, 2)
	graph.InsertOfA(2, 3)
	graph.InsertOfA(2, 4)
	graph.InsertOfA(2, 5)

	graph.Print()
}
