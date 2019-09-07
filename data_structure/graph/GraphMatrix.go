package graph
//邻接矩阵
type GraphMatrix struct {
	datas [][]int
	size int
}

func NewGraphMatrix(size int) *GraphMatrix{
	return &GraphMatrix{size:size}
}

func (graph *GraphMatrix) AddEdge(a int, b int, weight int) bool {
	if a > graph.size || b > graph.size {
		return false
	}

	return true
}

func (graph *GraphMatrix) RemoveEdge(a int, b int) bool  {
	if a > graph.size || b > graph.size {
		return false
	}
	return true
}
