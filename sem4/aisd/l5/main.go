package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strconv"
)

type Edge struct {
	from int
	to   int
	cost float64
}

type Item struct {
	value    Edge
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  //Avoid memory leak
	item.index = -1 // Negative index so you know something is wrong
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Edge, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

const (
	N = 5
)

type Graph struct {
	size    int
	matrix  [][]float64
	visited []bool
}

func newGraph(n int) *Graph {
	matrix := createMatrix(n)
	return &Graph{size: n, matrix: matrix, visited: make([]bool, n)}
}

func createMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}
	return matrix
}

func newRandomGraph(n int) *Graph {
	graph := newGraph(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			weight := rand.Float64()
			graph.matrix[i][j] = weight
			graph.matrix[j][i] = weight
		}
	}
	return graph
}

func (g *Graph) print() {
	for i := 0; i < g.size; i++ {
		fmt.Print("[ ")
		for j := 0; j < g.size; j++ {
			fmt.Printf("%.2f", g.matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Print("]\n")
	}
}

func lazyPrim(g *Graph) {
	node := 0
	g.visited[node] = true

	pq := &PriorityQueue{}
	heap.Init(pq)

	addEdges := func(currentNode int) {
		for adj := 0; adj < g.size; adj++ {
			if !g.visited[adj] && g.matrix[currentNode][adj] != 0 {
				heap.Push(pq, &Item{
					value: Edge{
						from: currentNode,
						to:   adj,
						cost: g.matrix[node][adj],
					},
					priority: -int(g.matrix[node][adj] * 100_000),
				})
			}
		}
	}

	addEdges(node) //Adding initial edges

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		edge := item.value

		if !g.visited[edge.to] {
			g.visited[edge.to] = true
			addEdges(edge.to)
			fmt.Println("Edge from:", strconv.Itoa(edge.from), "to:", strconv.Itoa(edge.to), "cost:", edge.cost)
		}

	}

}

func main() {
	fmt.Println("Hello Graphs!")

	graph := newRandomGraph(N)
	graph.print()

	lazyPrim(graph)

}
