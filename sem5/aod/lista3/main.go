package main

import (
	"dijkstra/dimacs"
	"dijkstra/graphs"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type DistanceResult struct {
	From, To, Distance int
}

func main() {

	algorithmName := flag.String("algo", "dijkstra", "Nazwa algorytmu (dijkstra, dial, radixheap)")
	dataFilePath := flag.String("d", "", "Ścieżka do pliku .gr z danymi (sieć + koszty)")
	sourcesFilePath := flag.String("ss", "", "Ścieżka do pliku .ss ze źródłami (opcjonalne)")
	outputSSFilePath := flag.String("oss", "", "Ścieżka do pliku wynikowego dla SS (opcjonalne)")

	pairsFilePath := flag.String("p2p", "", "Ścieżka do pliku .p2p z parami (opcjonalne)")
	outputP2PFilePath := flag.String("op2p", "", "Ścieżka do pliku wynikowego dla P2P (opcjonalne)")

	flag.Parse()

	// Sprawdzamy minimalne wymagania
	if *dataFilePath == "" {
		fmt.Println("Musisz podać przynajmniej flagę -d z plikiem danych (.gr).")
		flag.Usage()
		os.Exit(1)
	}

	dataFile, err := os.Open(*dataFilePath)
	if err != nil {
		log.Fatalf("Nie można otworzyć pliku z danymi: %v", err)
	}
	defer dataFile.Close()

	graph, err := dimacs.ParseFile(dataFile)
	if err != nil {
		log.Fatalf("Błąd parsowania pliku z danymi: %v", err)
	}
	switch {
	case *sourcesFilePath != "" && *outputSSFilePath != "":
		// Tryb "single-source"
		// Wczytujemy wszystkie źródła z pliku .ss
		sf, err := os.Open(*sourcesFilePath)
		if err != nil {
			log.Fatalf("Nie można otworzyć pliku ze źródłami: %v", err)
		}
		defer sf.Close()

		sources, err := dimacs.ParseSources(sf)
		if err != nil {
			log.Fatalf("Błąd parsowania źródeł: %v", err)
		}

		if len(sources) == 0 {
			log.Fatal("Brak źródeł w pliku .ss.")
		}

		var totalTime float64
		for _, s := range sources {
			start := time.Now()

			_, _ = runSelectedAlgorithm(graph, s, *algorithmName)

			elapsed := time.Since(start).Seconds() * 1000.0 // czas w ms
			totalTime += elapsed
		}
		avgTime := totalTime / float64(len(sources))

		// Tutaj - jeżeli potrzebujesz - można wypisać lub zapisać do pliku
		// wyniki poszczególnych algorytmów. Według opisu w zadaniu
		// wystarczy jednak, by do pliku .res trafił w formacie DIMACS
		// "t <average_time_in_msec>" + info o plikach itp.

		// Na potrzeby raportu do pliku .res generujemy krótkie podsumowanie
		// (w tym: nazwa pliku z danymi, plik ze źródłami, wymiary grafu, min/max cost, itp.)
		// Załóżmy, że w razie potrzeby potrafisz wydobyć liczbę wierzchołków, łuków
		// i min/max koszt z samego grafu (np. graph.Size(), graph.EdgeCount(), graph.MinCost, graph.MaxCost)
		size := graph.Size
		edgeCount := graph.EdgeCount
		maxCost, minCost := graphs.MaxEdgeCost(graph), graphs.MinEdgeCost(graph)

		resContent := dimacs.GenerateSSResults(
			*dataFilePath,
			*sourcesFilePath,
			time.Duration(avgTime)*time.Millisecond,
			size,
			edgeCount,
			maxCost,
			minCost,
		)

		if err := os.WriteFile(*outputSSFilePath, []byte(resContent), 0644); err != nil {
			log.Fatalf("Nie można zapisać wyników do pliku: %v", err)
		}

	case *pairsFilePath != "" && *outputP2PFilePath != "":
		// Tryb "point-to-point"
		// Wczytujemy wszystkie pary z pliku .p2p
		pf, err := os.Open(*pairsFilePath)
		if err != nil {
			log.Fatalf("Nie można otworzyć pliku z parami: %v", err)
		}
		defer pf.Close()

		pairs, err := dimacs.ParsePairs(pf)
		if err != nil {
			log.Fatalf("Błąd parsowania pliku z parami: %v", err)
		}

		// W tym miejscu wyznaczamy najkrótsze ścieżki między każdą z par (From, To).
		// W zależności od wybranego algorytmu (dijkstra/dial/radixheap),
		// można albo wyliczyć każdą parę niezależnie, albo wykorzystać
		// najpierw single-source i czytać wyniki. To już implementacja dowolna.
		//
		// Poniżej tylko szkic:

		var distances []DistanceResult

		for _, pair := range pairs {
			// Przykład:
			// distMap := runSelectedAlgorithm(graph, pair.From, algorithmName)
			// d := distMap[pair.To]
			d, _ := runSelectedAlgorithm(graph, pair.From, *algorithmName)
			distance := d[pair.To]

			distances = append(distances, DistanceResult{
				From:     pair.From,
				To:       pair.To,
				Distance: distance,
			})
		}

		size := graph.Size
		edgeCount := graph.EdgeCount
		maxCost, minCost := graphs.MaxEdgeCost(graph), graphs.MinEdgeCost(graph)

		output := dimacs.GenerateP2PResults(
			*dataFilePath,
			*pairsFilePath,
			size,
			edgeCount,
			maxCost,
			minCost,
			toDimacsFormat(distances),
		)

		if err := os.WriteFile(*outputP2PFilePath, []byte(output), 0644); err != nil {
			log.Fatalf("Nie można zapisać wyników do pliku: %v", err)
		}

	default:
		fmt.Println("Niepoprawne wywołanie programu.")
		fmt.Println("Musisz podać jednocześnie -ss i -oss LUB -p2p i -op2p.")
		flag.Usage()
		os.Exit(1)
	}
}

func runSelectedAlgorithm(g *graphs.Graph, source int, algoName string) ([]int, error) {
	switch algoName {
	case "dijkstra":
		dist, _, err := graphs.DijkstraBasic(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	case "dial":
		dist, _, err := graphs.DijkstraDial(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	case "radixheap":
		dist, _, err := graphs.DijkstraRadix(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	default:
		dist, _, err := graphs.DijkstraBasic(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	}
}

func toDimacsFormat(results []DistanceResult) []struct {
	From, To, Distance int
} {
	out := make([]struct{ From, To, Distance int }, 0, len(results))

	for _, r := range results {
		// Jeżeli w twoim grafie wewnętrznie przechowujesz wierzchołki 0-based,
		// a w formacie DIMACS chcesz mieć je 1-based, wystarczy dodać 1.
		out = append(out, struct {
			From, To, Distance int
		}{
			From:     r.From + 1, // 0-based => 1-based
			To:       r.To + 1,   // 0-based => 1-based
			Distance: r.Distance, // sam koszt się nie zmienia
		})
	}

	return out
}
