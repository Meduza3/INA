package graph

import (
	"fmt"
	"hamming/hamming"
	"math"
)

// AdjacencyEdge holds information for an edge in the adjacency list.
type AdjacencyEdge struct {
	to       int // endpoint of this edge
	capacity int // capacity of this edge
	flow     int // current flow on this edge
	rev      int // index of the reverse edge in adjacency[to]
}

// AdjacencyListGraph holds the flow network in adjacency list form.
type AdjacencyListGraph struct {
	size    int               // number of nodes
	adj     [][]AdjacencyEdge // adjacency list; adj[u] holds edges out of u
	maxFlow int               // track the maximum flow after running Edmonds-Karp
}

// NewAdjacencyListGraph creates a new graph of size n with empty adjacency lists.
func NewAdjacencyListGraph(n int) *AdjacencyListGraph {
	return &AdjacencyListGraph{
		size: n,
		adj:  make([][]AdjacencyEdge, n),
	}
}

func NewHypercubeAdjacencyList(k int) *AdjacencyListGraph {
	n := int(math.Pow(2, float64(k)))
	g := NewAdjacencyListGraph(n)

	for from := 0; from < n; from++ {
		for to := 0; to < n; to++ {
			if from == to {
				continue
			}

			fw := hamming.GetHammingWeight(from)
			tw := hamming.GetHammingWeight(to)
			if fw < tw && hamming.GetHammingDistance(from, to) == 1 {
				cap := generateCapacity(from, to)
				g.AddEdge(from, to, cap)
			}
		}
	}

	return g
}

func (g *AdjacencyListGraph) PrintFlowEdges() {
	fmt.Println("=== Flow on edges ===")
	for u := 0; u < g.size; u++ {
		for _, edge := range g.adj[u] {
			// Jeśli chcesz wyświetlać tylko krawędzie z dodatnim przepływem
			// możesz użyć warunku edge.flow > 0.
			// Jeśli chcesz wszystkie krawędzie (nawet flow == 0), usuń warunek.
			if edge.capacity > 0 { // krawędź 'do przodu'
				fmt.Printf("Flow(%d -> %d) = %d\n", u, edge.to, edge.flow)
			}
		}
	}
}

// AddEdge adds a directed edge (u -> v) of capacity cap, along with the reverse edge (v -> u) of capacity 0.
func (g *AdjacencyListGraph) AddEdge(u, v, cap int) {
	// forward edge
	fwd := AdjacencyEdge{
		to:       v,
		capacity: cap,
		flow:     0,
		rev:      len(g.adj[v]), // index of the reverse edge in g.adj[v]
	}
	// reverse edge
	rev := AdjacencyEdge{
		to:       u,
		capacity: 0,
		flow:     0,
		rev:      len(g.adj[u]), // index of the forward edge in g.adj[u]
	}

	g.adj[u] = append(g.adj[u], fwd)
	g.adj[v] = append(g.adj[v], rev)
}

// EdmondsKarp runs the Edmonds-Karp algorithm to find the maximum flow from source to sink.
func (g *AdjacencyListGraph) EdmondsKarp(source, sink int) int {
	parent := make([]int, g.size)     // parent[v] will store the node that leads to v
	parentEdge := make([]int, g.size) // parentEdge[v] will store which edge index in adj[parent[v]] leads to v

	maxFlow := 0

	// While we can find a path with available capacity in the residual graph
	for g.bfs(source, sink, parent, parentEdge) {
		// Find the minimum residual capacity along this path
		pathFlow := math.MaxInt
		v := sink
		for v != source {
			u := parent[v]
			eIndex := parentEdge[v]
			edge := &g.adj[u][eIndex]
			residual := edge.capacity - edge.flow
			if residual < pathFlow {
				pathFlow = residual
			}
			v = u
		}

		// Update flows along this path
		v = sink
		for v != source {
			u := parent[v]
			eIndex := parentEdge[v]
			revIndex := g.adj[u][eIndex].rev

			// Update forward edge
			g.adj[u][eIndex].flow += pathFlow
			// Update reverse edge
			g.adj[v][revIndex].flow -= pathFlow

			v = u
		}
		maxFlow += pathFlow
	}

	g.maxFlow = maxFlow
	return maxFlow
}

// bfs finds a path from source to sink using available (residual) capacities.
// It fills out 'parent' and 'parentEdge' arrays if it finds a path. Returns true if found.
func (g *AdjacencyListGraph) bfs(source, sink int, parent, parentEdge []int) bool {
	visited := make([]bool, g.size)
	for i := range parent {
		parent[i] = -1
		parentEdge[i] = -1
	}

	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// Explore all edges out of u
		for i, edge := range g.adj[u] {
			v := edge.to
			// If not visited and the residual capacity > 0
			if !visited[v] && edge.capacity-edge.flow > 0 {
				visited[v] = true
				parent[v] = u
				parentEdge[v] = i

				// If we've reached the sink, stop BFS
				if v == sink {
					return true
				}
				queue = append(queue, v)
			}
		}
	}

	return false
}

// PrintGraph prints the edges of the flow network: capacity and current flow.
func (g *AdjacencyListGraph) PrintGraph() {
	for u := 0; u < g.size; u++ {
		for _, edge := range g.adj[u] {
			// Show only the forward edges with positive capacity
			if edge.capacity > 0 {
				fmt.Printf("%d -> %d (capacity: %d, flow: %d)\n",
					u, edge.to, edge.capacity, edge.flow)
			}
		}
	}
}
