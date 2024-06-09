package main

import (
	"fmt"
	"math/rand"
)

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

func main() {
	fmt.Println("Hello Graphs!")

	graph := newRandomGraph(N)
	printGraph(graph)

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
			weight = rand.Float64()
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
