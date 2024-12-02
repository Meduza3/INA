package graph

type Graph struct {
	matrix [][]int
}

func NewGraph(size int) Graph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return Graph{
		matrix: matrix,
	}
}

func (g *Graph) AddEdge(from, to, cost int) {
	g.matrix[from][to] = cost
}
