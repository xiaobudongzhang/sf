package graph

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)
func TestGraphList_AddEdge(t *testing.T) {
	lib.PrintFunc()
	graph := NewGraphList(10)
	graph.AddEdge(1, 1)
	graph.AddEdge(1, 4)
	graph.Print()
}

func TestGraphList_RemoveEdge(t *testing.T) {
	lib.PrintFunc()
	lib.PrintFunc()
	graph := NewGraphList(10)
	graph.AddEdge(1, 1)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)
	graph.Print()
	graph.RemoveEdge(1, 4)
	graph.Print()
}
