package graphs

import (
	"fmt"
	"math"
)

type Graph struct {
	Size      int
	EdgeCount int
	data      []int
}

func NewGraph(size int) *Graph {
	data := make([]int, size*size)
	for i := range data {
		data[i] = -1
	}
	return &Graph{
		Size: size,
		data: data,
	}
}

func (g *Graph) AddEdge(from, to int, cost int) error {
	if from >= g.Size || to >= g.Size {
		return fmt.Errorf("out of bounds! size=%d from=%d to=%d", g.Size, from, to)
	}
	g.data[from*g.Size+to] = cost
	g.EdgeCount++
	return nil
}

func MaxEdgeCost(g *Graph) int {
	maxCost := 0
	for u := 0; u < g.Size; u++ {
		for v := 0; v < g.Size; v++ {
			cost := g.data[u*g.Size+v]
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

func MinEdgeCost(g *Graph) int {
	minCost := math.MaxInt
	for u := 0; u < g.Size; u++ {
		for v := 0; v < g.Size; v++ {
			cost := g.data[u*g.Size+v]
			if cost < 0 {
				continue
			}
			if cost < minCost {
				minCost = cost
			}
		}
	}
	return minCost
}
