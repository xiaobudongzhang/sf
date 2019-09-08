package graph

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestKahn_TopoSort(t *testing.T) {
	lib.PrintFunc()

	kahn := NewKahn(10)
	//kahn.graph.AddEdge(0, 1)
	//kahn.graph.AddEdge(0, 3)
	kahn.graph.AddEdge(1, 2)
	kahn.graph.AddEdge(1, 4)
	kahn.graph.AddEdge(2, 5)
	//kahn.graph.AddEdge(3, 4)
	kahn.graph.AddEdge(4, 5)
	//kahn.graph.AddEdge(4, 6)
	kahn.graph.AddEdge(5, 7)
	//kahn.graph.AddEdge(6, 7)

	kahn.graph.Print()


	fmt.Println("---------------------------------------------------------------------------")
	kahn.TopoSort()
}
