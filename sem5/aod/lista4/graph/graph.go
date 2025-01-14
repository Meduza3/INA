package graph

import (
	"fmt"
	"hamming/hamming"
	"math"
	"math/rand"
)

// Edge represents an edge in a flow network.
type Edge struct {
	To       int // The vertex this edge goes to
	Rev      int // Index of the reverse edge in adjList[To]
	Capacity int // How much flow this edge can carry
	Flow     int // Current flow in this edge
}

// Hypercube represents a flow network with adjacency list storage.
type Hypercube struct {
	hypercube_size int
	adjList        [][]Edge
}

func (g *Hypercube) PrintGraph() {
	for u, edges := range g.adjList {
		fmt.Printf("%s -> [", hamming.GetBit(u, g.hypercube_size))
		for i, e := range edges {
			if e.Capacity != 0 {
				fmt.Printf(" (To:%s,Cap:%d,Flow:%d,Rev:%d)", hamming.GetBit(e.To, g.hypercube_size), e.Capacity, e.Flow, e.Rev)
				if i < len(edges)-1 {
					fmt.Print(",")
				}
			}
		}
		fmt.Println(" ]")
	}
}

func NewHypercube(k int) *Hypercube {
	g := NewGraph(k)

	for from := range g.adjList {
		for to := range g.adjList {
			if from == to {
				continue
			}
			fw := hamming.GetHammingWeight(from)
			tw := hamming.GetHammingWeight(to)

			if fw < tw && hamming.GetHammingDistance(from, to) == 1 {
				capacity := generateCapacity(from, to)
				g.AddEdge(from, to, capacity)
			}
		}
	}
	return g
}

func generateCapacity(from, to int) int {
	h_from := hamming.GetHammingWeight(from)
	z_from := hamming.GetZeroWeight(from)
	h_to := hamming.GetHammingWeight(to)
	z_to := hamming.GetZeroWeight(to)

	max := max(h_from, z_from, h_to, z_to)
	two_l := int(math.Pow(2, float64(max)))
	capacity := rand.Intn(two_l) + 1
	return capacity
}

// NewGraph creates a graph with n vertices (0 through n-1).
func NewGraph(n int) *Hypercube {
	return &Hypercube{
		hypercube_size: n,
		adjList:        make([][]Edge, int(math.Pow(2, float64(n)))),
	}
}

// AddEdge adds a directed edge u -> v with the given capacity
// and automatically adds the reverse edge v -> u with zero capacity.
func (g *Hypercube) AddEdge(u, v, capacity int) {
	// Forward edge: from u to v
	fwdEdge := Edge{
		To:       v,
		Rev:      len(g.adjList[v]), // index where reverse edge will be added
		Capacity: capacity,
		Flow:     0,
	}
	// Reverse edge: from v to u, capacity 0 initially
	revEdge := Edge{
		To:       u,
		Rev:      len(g.adjList[u]), // index where forward edge is
		Capacity: 0,
		Flow:     0,
	}

	g.adjList[u] = append(g.adjList[u], fwdEdge)
	g.adjList[v] = append(g.adjList[v], revEdge)
}

func FindShortestPath(g *Hypercube, from, to int) {

}

// 1. Run BFS on edges

func (e *Edge) isCapable() bool {
	return e.Flow < e.Capacity
}

func edgeCountFromSize(hypercube_size int) int {
	return hypercube_size * int(math.Pow(2, float64(hypercube_size)-1))
}

// func BFS(g *Hypercube, source, sink int) []int {
// 	queue := []int{}
// 	for _, edge := range g.adjList[source] {
// 		if edge.isCapable() {
// 			if edge.To == 1 {
// 				queue = append(queue, edge.To)
// 			}
// 		}

// 		if len(queue) == 0 {
// 			return nil // No paths can come
// 		}

// 		for len(queue) > 0 {
// 			current := queue[0]
// 			queue = queue[1:]

// 			for _, edge := range g.adjList[current] {
// 				if edge.isCapable() {
// 					queue = append(queue, edge.To)
// 				}
// 			}
// 		}
// 	}
// }
