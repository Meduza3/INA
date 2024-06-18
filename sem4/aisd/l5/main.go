package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type MSTree struct {
	AdjList [][]int
}

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
	N = 1000
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

func find(node int, ufs []int) int {
	if ufs[node] != node {
		ufs[node] = find(ufs[node], ufs)
	}
	return ufs[node]
}

func kruskal(g *Graph) []Edge {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < g.size; i++ {
		for j := i + 1; j < g.size; j++ {
			if g.matrix[i][j] != 0 {
				heap.Push(pq, &Item{
					value:    Edge{from: i, to: j, cost: g.matrix[i][j]},
					priority: -int(g.matrix[i][j] * 100_000),
				})
			}
		}
	}

	ufs := make([]int, g.size)
	for i := range ufs {
		ufs[i] = i
	}

	same := func(node1, node2 int) bool {
		return find(node1, ufs) == find(node2, ufs)
	}

	union := func(node1, node2 int) {
		ufs[find(node1, ufs)] = find(node2, ufs)
	}

	mst := []Edge{}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		edge := item.value

		if !same(edge.from, edge.to) {
			union(edge.from, edge.to)
			mst = append(mst, Edge{from: edge.from, to: edge.to, cost: edge.cost})
		}
	}

	return mst
}

func prim(g *Graph) []Edge {
	node := 0
	g.visited[node] = true

	pq := &PriorityQueue{}
	heap.Init(pq)

	mst := []Edge{}

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
			mst = append(mst, Edge{from: edge.from, to: edge.to, cost: edge.cost})
		}

	}

	return mst
}

func main() {
	fmt.Println("Hello Graphs!")

	file, err := os.Create("output.csv")
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"n", "timePrim", "timeKruskal"})

	for n := 10; n <= 500; n += 10 {
		graph := newRandomGraph(n)
		//graph.print()

		startTimePrim := time.Now()
		_ = prim(graph)
		timeTakenPrim := time.Since(startTimePrim).Milliseconds()

		startTimeKruskal := time.Now()
		_ = kruskal(graph)
		timeTakenKruskal := time.Since(startTimeKruskal).Milliseconds()

		fmt.Println(strconv.Itoa(n) + " : " + strconv.FormatInt(timeTakenPrim, 10) + " : " + strconv.FormatInt(timeTakenKruskal, 10))
		writer.Write([]string{strconv.Itoa(n), strconv.FormatInt(timeTakenPrim, 10), strconv.FormatInt(timeTakenKruskal, 10)})
	}

}
