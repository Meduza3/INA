// experiments/topological_sort_experiment.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"graphexperiment/graph"
)

// TopologicalSortResult przechowuje wynik sortowania topologicznego
type TopologicalSortResult struct {
	HasCycle  bool
	Order     []int
	TimeTaken time.Duration
}

// TopologicalSort wykonuje sortowanie topologiczne za pomocą algorytmu Kahn'a
func TopologicalSort(g *graph.Graph) (TopologicalSortResult, error) {
	start := time.Now()
	inDegree := make([]int, g.N+1)
	for u := 1; u <= g.N; u++ {
		for _, v := range g.Adj[u] {
			inDegree[v]++
		}
	}

	queue := []int{}
	for u := 1; u <= g.N; u++ {
		if inDegree[u] == 0 {
			queue = append(queue, u)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)

		for _, v := range g.Adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	elapsed := time.Since(start)

	if len(order) != g.N {
		return TopologicalSortResult{
			HasCycle:  true,
			Order:     nil,
			TimeTaken: elapsed,
		}, nil
	}

	return TopologicalSortResult{
		HasCycle:  false,
		Order:     order,
		TimeTaken: elapsed,
	}, nil
}

// WriteResult zapisuje wynik sortowania do pliku
func WriteResult(filename string, result TopologicalSortResult, g *graph.Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if result.HasCycle {
		writer.WriteString("Graf zawiera cykl skierowany.\n")
	} else {
		writer.WriteString("Graf jest acykliczny.\n")
		if g.N <= 200 {
			writer.WriteString("Porządek topologiczny wierzchołków:\n")
			for i, v := range result.Order {
				writer.WriteString(fmt.Sprintf("%d", v))
				if i != len(result.Order)-1 {
					writer.WriteString(" ")
				}
			}
			writer.WriteString("\n")
		}
	}
	writer.WriteString(fmt.Sprintf("Czas wykonania programu: %v\n", result.TimeTaken))

	return nil
}

func main() {
	// Definiowanie flag
	dirPtr := flag.String("dir", "./test_2", "Ścieżka do katalogu z plikami grafów")
	outputPtr := flag.String("output", "./results_topo", "Ścieżka do katalogu, gdzie zostaną zapisane wyniki")
	flag.Parse()

	// Tworzenie katalogu wyników, jeśli nie istnieje
	if _, err := os.Stat(*outputPtr); os.IsNotExist(err) {
		err := os.MkdirAll(*outputPtr, os.ModePerm)
		if err != nil {
			fmt.Printf("Błąd podczas tworzenia katalogu wyników: %v\n", err)
			os.Exit(1)
		}
	}

	// Listowanie plików w katalogu
	files, err := ioutil.ReadDir(*dirPtr)
	if err != nil {
		fmt.Printf("Błąd podczas odczytu katalogu %s: %v\n", *dirPtr, err)
		os.Exit(1)
	}

	// Iteracja po każdym pliku
	for _, file := range files {
		if file.IsDir() {
			continue // pomijanie podkatalogów
		}

		filePath := filepath.Join(*dirPtr, file.Name())
		fmt.Printf("Przetwarzanie pliku: %s\n", filePath)

		// Wczytanie grafu
		g, err := graph.ReadGraph(filePath)
		if err != nil {
			fmt.Printf("  Błąd wczytywania grafu: %v\n", err)
			continue
		}

		// Wykonanie sortowania topologicznego
		result, err := TopologicalSort(g)
		if err != nil {
			fmt.Printf("  Błąd podczas sortowania topologicznego: %v\n", err)
			continue
		}

		// Przygotowanie nazwy pliku wynikowego
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		resultFileName := fmt.Sprintf("%s_topo.txt", baseName)
		resultFilePath := filepath.Join(*outputPtr, resultFileName)

		// Zapisanie wyniku do pliku
		err = WriteResult(resultFilePath, result, g)
		if err != nil {
			fmt.Printf("  Błąd zapisu wyniku: %v\n", err)
			continue
		}

		fmt.Printf("  Wynik zapisany do: %s\n", resultFilePath)
	}

	fmt.Println("Eksperymenty zakończone. Wyniki zapisane w folderze:", *outputPtr)
}
