package main

import (
	"dijkstra/dimacs"
	"dijkstra/graphs"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define command-line flags
	dataFile := flag.String("d", "", "Path to the DIMACS data file")
	//sourceFile := flag.String("ss", "", "Path to the source vertices file")
	//outputFile := flag.String("oss", "", "Path to the output results file")
	flag.Parse()

	if *dataFile == "" /*|| *sourceFile == "" || *outputFile == "" */ {
		fmt.Println("Usage: dijkstra -d <data file> -ss <source file> -oss <output file>")
		os.Exit(1)
	}

	data, err := os.Open(*dataFile)
	if err != nil {
		fmt.Printf("Error opening data file: %v\n", err)
		os.Exit(1)
	}
	defer data.Close()

	graph, err := dimacs.ParseFile(data)
	if err != nil {
		fmt.Printf("Error parsing data file: %v\n", err)
		os.Exit(1)
	}

	dist, parent, err := graphs.DijkstraBasic(graph, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Djikstra: %+v ||| %+v\n", dist, parent)
	distDial, parentDial, err := graphs.DijkstraDial(graph, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Dial    : %+v ||| %+v\n", distDial, parentDial)

	distRadix, parentRadix, err := graphs.DijkstraRadix(graph, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Radix   : %+v ||| %+v\n", distRadix, parentRadix)
}
