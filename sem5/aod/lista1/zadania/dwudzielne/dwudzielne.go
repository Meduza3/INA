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

// BipartiteResult przechowuje wynik sprawdzania dwudzielności
type BipartiteResult struct {
	IsBipartite bool
	Partition   [2][]int
	TimeTaken   time.Duration
}

// WriteResult zapisuje wynik sprawdzania dwudzielności do pliku
func WriteResult(filename string, result BipartiteResult, g *graph.Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	if result.IsBipartite {
		writer.WriteString("Graf jest dwudzielny.\n")
		if g.N <= 200 {
			// Sortowanie dla czytelności
			for i := 0; i < 2; i++ {
				sort.Ints(result.Partition[i])
				writer.WriteString(fmt.Sprintf("Podzbiór V%d: ", i))
				for j, v := range result.Partition[i] {
					writer.WriteString(fmt.Sprintf("%d", v))
					if j != len(result.Partition[i])-1 {
						writer.WriteString(" ")
					}
				}
				writer.WriteString("\n")
			}
		}
	} else {
		writer.WriteString("Graf nie jest dwudzielny.\n")
	}

	writer.WriteString(fmt.Sprintf("Czas wykonania algorytmu sprawdzania dwudzielności: %v\n", result.TimeTaken))

	return nil
}

func main() {
	// Definiowanie flag
	dirPtr := flag.String("dir", "./test_4", "Ścieżka do katalogu z plikami grafów")
	outputPtr := flag.String("output", "./results_dwudzielne", "Ścieżka do katalogu, gdzie zostaną zapisane wyniki")
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

		// Wykonanie algorytmu sprawdzania dwudzielności
		start := time.Now()
		bipartite := g.IsBipartite()
		duration := time.Since(start)

		result := BipartiteResult{
			IsBipartite: bipartite.IsBipartite,
			Partition:   bipartite.Partition,
			TimeTaken:   duration,
		}

		// Przygotowanie nazwy pliku wynikowego
		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		resultFileName := fmt.Sprintf("%s_bipartite.txt", baseName)
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
