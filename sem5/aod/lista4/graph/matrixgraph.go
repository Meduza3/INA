package graph

import (
	"fmt"
	"hamming/hamming"
	"math"
)

type MatrixGraph struct {
	size     int // number of nodes (should be 2^k for a hypercube)
	exists   [][]bool
	capacity [][]int // capacity[u][v] is the capacity of edge (u -> v)
	flow     [][]int // flow[u][v] is the current flow of edge (u -> v)
	maxflow  int
}

func NewMatrixGraph(n int) *MatrixGraph {
	exists := make([][]bool, n)
	capacity := make([][]int, n)
	flow := make([][]int, n)
	for i := 0; i < n; i++ {
		exists[i] = make([]bool, n)
		capacity[i] = make([]int, n)
		flow[i] = make([]int, n)
	}

	return &MatrixGraph{
		size:     n,
		capacity: capacity,
		flow:     flow,
		exists:   exists,
	}
}

func (g *MatrixGraph) AddEdge(u, v, cap int) {
	g.exists[u][v] = true
	g.exists[v][u] = true
	g.capacity[u][v] = cap
	g.capacity[v][u] = 0
}

func NewHypercubeMatrix(k int) *MatrixGraph {
	n := int(math.Pow(2, float64(k)))
	mg := NewMatrixGraph(n)

	// Populate the matrix with capacities where edges exist
	for from := 0; from < n; from++ {
		for to := 0; to < n; to++ {
			if from == to {
				continue
			}

			fw := hamming.GetHammingWeight(from)
			tw := hamming.GetHammingWeight(to)
			if fw < tw && hamming.GetHammingDistance(from, to) == 1 {
				capacity := generateCapacity(from, to)
				mg.AddEdge(from, to, capacity)
			}
		}
	}
	return mg
}

func (g *MatrixGraph) PrintMatrix() {
	fmt.Println("Capacity Matrix:")

	fmt.Printf("     ")
	for v := 0; v < g.size; v++ {
		fmt.Printf("%3d", v)
	}
	fmt.Println()

	// Print a separator line
	fmt.Printf("      ")
	for v := 0; v < g.size; v++ {
		fmt.Print("---")
	}
	fmt.Println()

	for u := 0; u < g.size; u++ {
		fmt.Printf("%3d: |", u)
		for v := 0; v < g.size; v++ {
			if g.exists[u][v] {
				fmt.Printf("%3d", g.capacity[u][v])
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
}

func (g *MatrixGraph) EdmondsKarp(source, sink int) int {

	for {
		path := g.BFS(source, sink)
		if path == nil {
			break
		}
		var minCapacity int = math.MaxInt
		for i := 0; i < len(path)-1; i++ {
			remainingCapacity := g.capacity[path[i]][path[i+1]] - g.flow[path[i]][path[i+1]]
			if remainingCapacity < minCapacity {
				minCapacity = remainingCapacity
			}
		}

		for i := 0; i < len(path)-1; i++ {
			g.flow[path[i]][path[i+1]] += minCapacity
		}
		g.maxflow += minCapacity
	}

	return g.maxflow
}

func (g *MatrixGraph) BFS(source, sink int) []int {
	visited := make([]bool, g.size)
	parent := make([]int, g.size)
	for i := range parent {
		parent[i] = -1
	}

	queue := []int{source}
	visited[source] = true

	foundSink := false
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 0; v < g.size; v++ {
			if g.exists[u][v] {
				if !visited[v] && g.capacity[u][v]-g.flow[u][v] > 0 {
					queue = append(queue, v)
					visited[v] = true
					parent[v] = u
					if v == sink {
						foundSink = true
						break
					}
				}
			}
			if foundSink {
				break
			}
		}
	}
	if !foundSink {
		return nil
	}

	var path []int
	for cur := sink; cur != -1; cur = parent[cur] {
		path = append([]int{cur}, path...)
	}
	fmt.Printf("path: %v\n", path)
	return path
}
