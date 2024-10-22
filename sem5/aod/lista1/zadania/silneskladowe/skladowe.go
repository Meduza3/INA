package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"graphexperiment/graph"
)

// SCCResult przechowuje wynik algorytmu SCC
type SCCResult struct {
	Components [][]int
	TimeTaken  time.Duration
}

// WriteResult zapisuje wynik SCC do pliku
func WriteResult(filename string, result SCCResult, g *graph.Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	numComponents := len(result.Components)
	writer.WriteString(fmt.Sprintf("Liczba silnie spójnych składowych: %d\n", numComponents))

	for i, component := range result.Components {
		writer.WriteString(fmt.Sprintf("Składowa %d: %d wierzchołków\n", i+1, len(component)))
		if g.N <= 200 {
			// Sortowanie wierzchołków w składowej dla czytelności
			sort.Ints(component)
			writer.WriteString("Wierzchołki: ")
			for j, v := range component {
				writer.WriteString(fmt.Sprintf("%d", v))
				if j != len(component)-1 {
					writer.WriteString(" ")
				}
			}
			writer.WriteString("\n")
		}
	}

	writer.WriteString(fmt.Sprintf("Czas wykonania algorytmu SCC: %v\n", result.TimeTaken))

	return nil
}

func main() {
	// Definiowanie flag
	dirPtr := flag.String("dir", "./test_3", "Ścieżka do katalogu z plikami grafów")
	outputPtr := flag.String("output", "./results_skladowe", "Ścieżka do katalogu, gdzie zostaną zapisane wyniki")
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

		// Sprawdzenie, czy graf jest skierowany
		if !g.Directed {
			fmt.Printf("  Pomijanie nieskierowanego grafu: %s\n", file.Name())
			continue
		}

		// Wykonanie algorytmu SCC
		start := time.Now()
		scc := g.SCC()
		duration := time.Since(start)

		result := SCCResult{
			Components: scc.Components,
			TimeTaken:  duration,
		}

		// Przygotowanie nazwy pliku wynikowego
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		resultFileName := fmt.Sprintf("%s_scc.txt", baseName)
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
