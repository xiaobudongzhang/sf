package graph

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestDfs_TopoSort(t *testing.T) {
	lib.PrintFunc()

	dfs := NewDfs(10)
	dfs.graph.AddEdge(1, 2)
	dfs.graph.AddEdge(1, 4)
	dfs.graph.AddEdge(2, 5)
	dfs.graph.AddEdge(4, 5)
	dfs.graph.AddEdge(5, 7)

	dfs.graph.Print()


	fmt.Println("---------------------------------------------------------------------------")
	dfs.TopoSort()
}

