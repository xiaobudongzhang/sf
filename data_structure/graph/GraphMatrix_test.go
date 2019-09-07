package graph

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestGraphMatrix_AddEdge(t *testing.T) {
	lib.PrintFunc()
	graph := NewGraphMatrix(10)
	graph.AddEdge(1,1 ,2)
	fmt.Printf("graph:%v\n", graph.Data())
}

func TestGraphMatrix_RemoveEdge(t *testing.T) {
	lib.PrintFunc()
	graph := NewGraphMatrix(10)
	graph.AddEdge(1,1 ,2)
	graph.RemoveEdge(1,1)
	fmt.Printf("graph:%v\n", graph.Data())
}
