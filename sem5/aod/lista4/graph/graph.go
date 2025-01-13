package graph

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Edge represents an edge in a flow network.
type Edge struct {
	To       int // The vertex this edge goes to
	Rev      int // Index of the reverse edge in adjList[To]
	Capacity int // How much flow this edge can still carry
	Flow     int // Current flow in this edge
}

// Graph represents a flow network with adjacency list storage.
type Graph struct {
	size    int
	adjList [][]Edge
}

func (g *Graph) PrintGraph() {
	for u, edges := range g.adjList {
		fmt.Printf("%s -> [", GetBit(u, g.size))
		for i, e := range edges {
			if e.Capacity != 0 {
				fmt.Printf(" (To:%s,Cap:%d,Flow:%d,Rev:%d)", GetBit(e.To, g.size), e.Capacity, e.Flow, e.Rev)
				if i < len(edges)-1 {
					fmt.Print(",")
				}
			}
		}
		fmt.Println(" ]")
	}
}

func NewHypercube(k int) *Graph {
	g := NewGraph(k)

	for from := range g.adjList {
		for to := range g.adjList {
			if from == to {
				continue
			}
			fw := GetHammingWeight(from)
			tw := GetHammingWeight(to)

			if fw < tw && GetHammingDistance(from, to) == 1 {
				capacity := generateCapacity(from, to)
				g.AddEdge(from, to, capacity)
			}
		}
	}
	return g
}

func generateCapacity(from, to int) int {
	h_from := GetHammingWeight(from)
	z_from := GetZeroWeight(from)
	h_to := GetHammingWeight(to)
	z_to := GetZeroWeight(to)

	max := max(h_from, z_from, h_to, z_to)
	two_l := int(math.Pow(2, float64(max)))
	capacity := rand.Intn(two_l) + 1
	return capacity
}

// NewGraph creates a graph with n vertices (0 through n-1).
func NewGraph(n int) *Graph {
	return &Graph{
		size:    n,
		adjList: make([][]Edge, int(math.Pow(2, float64(n)))),
	}
}

// AddEdge adds a directed edge u -> v with the given capacity
// and automatically adds the reverse edge v -> u with zero capacity.
func (g *Graph) AddEdge(u, v, capacity int) {
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

func GetGreaterHamming(a, b *int) (greater, lesser *int) {
	ah := GetHammingWeight(*a)
	bh := GetHammingWeight(*b)
	if ah >= bh {
		return a, b
	} else {
		return b, a
	}
}

func GetHammingDistance(from, to int) int {
	fromBitstring := GetBit(from, -1)
	toBitstring := GetBit(to, -1)
	var distance int
	for i := 0; i < len(fromBitstring) || i < len(toBitstring); i++ {
		var fromBit, toBit byte = '0', '0'
		if i < len(fromBitstring) {
			fromBit = fromBitstring[len(fromBitstring)-1-i]
		}
		if i < len(toBitstring) {
			toBit = toBitstring[len(toBitstring)-1-i]
		}
		if fromBit != toBit {
			distance++
		}
	}
	return distance
}

func GetHammingWeight(num int) int {
	bitstring := GetBit(num, -1)
	var result int
	for _, c := range bitstring {
		if c == '1' {
			result++
		}
	}
	return result
}

func GetZeroWeight(num int) int {
	bitstring := GetBit(num, -1)
	var result int
	for _, c := range bitstring {
		if c == '0' {
			result++
		}
	}
	return result
}

func GetBit(num, length int) string {
	result := strconv.FormatInt(int64(num), 2)
	if length == -1 {
		return result
	}
	for len(result) < length {
		result = "0" + result
	}
	return result
}
