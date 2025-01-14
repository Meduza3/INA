package main

import (
	"flag"
	"fmt"
	"hamming/graph"
	"math"
)

func main() {
	size := flag.Int("size", 3, "size of the hypercube")
	printFlow := flag.Bool("printFlow", false, "do you want to print flow?")
	flag.Parse()
	g := graph.NewHypercubeAdjacencyList(*size)
	if *size <= 5 {
		g.PrintGraph()
	}
	sink := int(math.Pow(2, float64(*size))) - 1
	fmt.Printf("max flow: %d\n", g.EdmondsKarp(0, sink))

	if *printFlow {
		g.PrintFlowEdges()
	}
}
