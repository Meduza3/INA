package graphs

import (
	"fmt"
)

type Graph struct {
	size int
	data []int
}

func NewGraph(size int) *Graph {
	data := make([]int, size*size)
	for i := range data {
		data[i] = -1
	}
	return &Graph{
		size: size,
		data: data,
	}
}

func (g *Graph) AddEdge(from, to int, cost int) error {
	if from >= g.size || to >= g.size {
		return fmt.Errorf("out of bounds! size=%d from=%d to=%d", g.size, from, to)
	}
	g.data[from*g.size+to] = cost
	return nil
}

func maxEdgeCost(g *Graph) int {
	maxCost := 0
	for u := 0; u < g.size; u++ {
		for v := 0; v < g.size; v++ {
			cost := g.data[u*g.size+v]
			if cost < 0 {
				continue
			}
			if cost > maxCost {
				maxCost = cost
			}
		}
	}
	return maxCost
}
