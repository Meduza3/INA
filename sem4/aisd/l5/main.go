package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type Item struct {
	value    float64
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

func (pq *PriorityQueue) update(item *Item, value float64, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

const (
	N = 10
)

type Graph struct {
	size   int
	matrix [][]float64
}

func newGraph(n int) *Graph {
	matrix := createMatrix(n)
	return &Graph{size: n, matrix: matrix}
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

func printGraph(g *Graph) {
	for i := 0; i < g.size; i++ {
		fmt.Print("[ ")
		for j := 0; j < g.size; j++ {
			fmt.Printf("%.2f", g.matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Print("]\n")
	}
}

func main() {
	fmt.Println("Hello Graphs!")

	graph := newRandomGraph(N)
	printGraph(graph)

}
