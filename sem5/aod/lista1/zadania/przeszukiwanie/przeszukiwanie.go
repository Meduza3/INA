package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"graphexperiment/graph"
)

// PrintTraversalOrder prints the order of visited vertices
func PrintTraversalOrder(order []int) string {
	var sb strings.Builder
	sb.WriteString("Kolejność odwiedzania wierzchołków:\n")
	for i, v := range order {
		sb.WriteString(fmt.Sprintf("%d", v))
		if i != len(order)-1 {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("\n")
	return sb.String()
}

func PrintTraversalTree(parent []int) string {
	var sb strings.Builder
	sb.WriteString("Drzewo przeszukiwania (wierzchołek: rodzic):\n")
	for v := 1; v < len(parent); v++ {
		if parent[v] != -1 && parent[v] != 0 {
			sb.WriteString(fmt.Sprintf("%d: %d\n", v, parent[v]))
		} else if parent[v] == -1 {
			sb.WriteString(fmt.Sprintf("%d: nil\n", v))
		}
	}
	return sb.String()
}

func main() {
	// Define command-line flags
	dirPtr := flag.String("dir", "./test_1", "Path to the directory containing graph files")
	outputPtr := flag.String("output", "./results", "Path to the directory where results will be saved")
	flag.Parse()

	// Create output directory if it doesn't exist
	if _, err := os.Stat(*outputPtr); os.IsNotExist(err) {
		err := os.MkdirAll(*outputPtr, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// List all files in the directory
	files, err := ioutil.ReadDir(*dirPtr)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", *dirPtr, err)
		os.Exit(1)
	}

	// Iterate over each file
	for _, file := range files {
		if file.IsDir() {
			continue // skip subdirectories
		}

		filePath := filepath.Join(*dirPtr, file.Name())
		fmt.Printf("Przetwarzanie pliku: %s\n", filePath)

		// Read the graph
		graph, err := graph.ReadGraph(filePath)
		if err != nil {
			fmt.Printf("  Błąd wczytywania grafu: %v\n", err)
			continue
		}

		// Perform BFS
		startBFS := time.Now()
		bfsOrder, bfsParent := graph.BFS(1)
		durationBFS := time.Since(startBFS)

		// Perform DFS
		startDFS := time.Now()
		dfsOrder, dfsParent := graph.DFS(1)
		durationDFS := time.Since(startDFS)

		// Prepare output content
		var bfsContent strings.Builder
		bfsContent.WriteString(fmt.Sprintf("Algorytm BFS dla pliku %s\n", file.Name()))
		bfsContent.WriteString(PrintTraversalOrder(bfsOrder))
		bfsContent.WriteString(PrintTraversalTree(bfsParent))
		bfsContent.WriteString(fmt.Sprintf("Czas wykonania BFS: %v\n", durationBFS))

		var dfsContent strings.Builder
		dfsContent.WriteString(fmt.Sprintf("Algorytm DFS dla pliku %s\n", file.Name()))
		dfsContent.WriteString(PrintTraversalOrder(dfsOrder))
		dfsContent.WriteString(PrintTraversalTree(dfsParent))
		dfsContent.WriteString(fmt.Sprintf("Czas wykonania DFS: %v\n", durationDFS))

		// Define output file names
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		bfsFileName := fmt.Sprintf("%s_bfs.txt", baseName)
		dfsFileName := fmt.Sprintf("%s_dfs.txt", baseName)

		bfsFilePath := filepath.Join(*outputPtr, bfsFileName)
		dfsFilePath := filepath.Join(*outputPtr, dfsFileName)

		// Write BFS results
		err = ioutil.WriteFile(bfsFilePath, []byte(bfsContent.String()), 0644)
		if err != nil {
			fmt.Printf("  Błąd zapisu wyników BFS: %v\n", err)
			continue
		}

		// Write DFS results
		err = ioutil.WriteFile(dfsFilePath, []byte(dfsContent.String()), 0644)
		if err != nil {
			fmt.Printf("  Błąd zapisu wyników DFS: %v\n", err)
			continue
		}

		fmt.Printf("  BFS zapisane do: %s\n", bfsFilePath)
		fmt.Printf("  DFS zapisane do: %s\n", dfsFilePath)
	}

	fmt.Println("Eksperymenty zakończone. Wyniki zapisane w folderze:", *outputPtr)
}
