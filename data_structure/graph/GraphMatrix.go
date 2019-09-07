package graph
//邻接矩阵
type GraphMatrix struct {
	datas [][]int
	size int
}

func NewGraphMatrix(size int) *GraphMatrix{

	datas := make([][]int, size+1)
	for k, _ :=range datas  {
		subdata := make([]int, size + 1)
		datas[k] = subdata
	}
	return &GraphMatrix{
		datas:datas,
		size:size+1,
	}
}

func (graph *GraphMatrix) AddEdge(a int, b int, weight int) bool {
	if a > graph.size || b > graph.size || a < 1 || b < 1 || weight < 1{
		return false
	}
	graph.datas[a][b] = weight
	return true
}

func (graph *GraphMatrix) RemoveEdge(a int, b int) bool  {
	if a > graph.size || b > graph.size || a < 1 || b < 1 {
		return false
	}
	graph.datas[a][b] = 0
	return true
}


func (graph *GraphMatrix) Data() [][]int  {
	return graph.datas
}
