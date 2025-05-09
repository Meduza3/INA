package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
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

func createAdjList(mst []Edge, size int) [][]int {
	adjList := make([][]int, size)
	for _, edge := range mst {
		adjList[edge.from] = append(adjList[edge.from], edge.to)
		adjList[edge.to] = append(adjList[edge.to], edge.from)
	}
	return adjList
}

// DFS
func calculateSubtreeSizes(node int, adjList [][]int, visited []bool, subtreeSizes []int) int {
	visited[node] = true
	subtreeSize := 1
	for _, child := range adjList[node] {
		if !visited[child] {
			subtreeSize += calculateSubtreeSizes(child, adjList, visited, subtreeSizes)
		}
	}
	subtreeSizes[node] = subtreeSize
	return subtreeSize
}

func broadcastOrder(node int, adjList [][]int, visited []bool, subtreeSizes []int, order []int) {
	visited[node] = true
	children := []int{}

	// Gather all unvisited children
	for _, child := range adjList[node] {
		if !visited[child] {
			children = append(children, child)
		}
	}

	// Debug: Print the current node and its children
	fmt.Printf("Current Node: %d, Children: %v\n", node, children)

	// Sort children based on the size of their subtree, largest first
	if len(children) > 0 {
		sort.Slice(children, func(i, j int) bool {
			// Debug: Print subtree sizes comparison
			fmt.Printf("Comparing subtree sizes: %d (%d) vs %d (%d)\n", children[i], subtreeSizes[children[i]], children[j], subtreeSizes[children[j]])
			return subtreeSizes[children[i]] > subtreeSizes[children[j]]
		})
	}

	// Recursively determine order for children
	for _, child := range children {
		order = append(order, child)
		broadcastOrder(child, adjList, visited, subtreeSizes, order)
	}
}

func simulateBroadcasts(nodeCount, simulations int) []int {
	roundsData := make([]int, 0, simulations)

	for i := 0; i < simulations; i++ {
		graph := newRandomGraph(nodeCount)
		mst := prim(graph)
		adjList := createAdjList(mst, nodeCount)

		startNode := rand.Intn(nodeCount)
		rounds := simulateBroadcastFromNode(adjList, startNode, nodeCount)
		roundsData = append(roundsData, rounds)
	}
	return roundsData
}

func simulateBroadcastFromNode(adjList [][]int, startNode, nodeCount int) int {
	subtreeSizes := make([]int, nodeCount)
	visited := make([]bool, nodeCount)

	calculateSubtreeSizes(startNode, adjList, visited, subtreeSizes)

	// Reset visited for the actual broadcast simulation
	for i := range visited {
		visited[i] = false
	}

	order := make([]int, 0, nodeCount)
	order = append(order, startNode)
	broadcastOrder(startNode, adjList, visited, subtreeSizes, order) // corrected line

	return simulateRounds(order, adjList)
}

func simulateRounds(order []int, adjList [][]int) int {
	rounds := 0
	queue := []int{order[0]}
	visited := make([]bool, len(adjList))
	visited[order[0]] = true

	for len(queue) > 0 {
		nextQueue := []int{}
		for _, node := range queue {
			for _, child := range adjList[node] {
				if !visited[child] {
					visited[child] = true
					nextQueue = append(nextQueue, child)
				}
			}
		}
		queue = nextQueue
		if len(queue) > 0 {
			rounds++
		}
	}
	return rounds
}

func analyzeRounds(roundsData []int) (float64, int, int) {
	var total int
	minRounds := int(^uint(0) >> 1) // max int
	maxRounds := 0

	for _, rounds := range roundsData {
		total += rounds
		if rounds < minRounds {
			minRounds = rounds
		}
		if rounds > maxRounds {
			maxRounds = rounds
		}
	}

	average := float64(total) / float64(len(roundsData))
	return average, minRounds, maxRounds
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

	//Experiment 1:
	//writer.Write([]string{"n", "timePrim", "timeKruskal"})
	//for n := 10; n <= 500; n += 10 {
	//	graph := newRandomGraph(n)
	//graph.print()

	//	startTimePrim := time.Now()
	//	_ = prim(graph)
	//	timeTakenPrim := time.Since(startTimePrim).Milliseconds()

	//	startTimeKruskal := time.Now()
	//	_ = kruskal(graph)
	//	timeTakenKruskal := time.Since(startTimeKruskal).Milliseconds()

	//	fmt.Println(strconv.Itoa(n) + " : " + strconv.FormatInt(timeTakenPrim, 10) + " : " + strconv.FormatInt(timeTakenKruskal, 10))
	//	writer.Write([]string{strconv.Itoa(n), strconv.FormatInt(timeTakenPrim, 10), strconv.FormatInt(timeTakenKruskal, 10)})
	//}

	//Experiment 2:
	err = writer.Write([]string{"Number of Nodes", "Average Rounds", "Minimum Rounds", "Maximum Rounds"})
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for n := 10; n <= 1000; n += 10 {
		roundsData := simulateBroadcasts(n, 100)
		avgRounds, minRounds, maxRounds := analyzeRounds(roundsData)
		err = writer.Write([]string{
			strconv.Itoa(n),
			fmt.Sprintf("%.2f", avgRounds),
			strconv.Itoa(minRounds),
			strconv.Itoa(maxRounds),
		})
		if err != nil {
			fmt.Println("Error writing data:", err)
			return
		}
		writer.Flush()
	}
}
