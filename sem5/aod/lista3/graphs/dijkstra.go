package graphs

import (
	"container/heap"
	"dijkstra/graphs/radix"
	"fmt"
	"math"
)

func DijkstraBasic(g *Graph, source int) ([]int, []int, error) {
	if source < 0 || source >= g.Size {
		return nil, nil, fmt.Errorf("source out of bounds: %d", source)
	}

	// Przygotowujemy tablicę odległości - startowo "nieskończoności".
	dist := make([]int, g.Size)
	parent := make([]int, g.Size)

	for i := range dist {
		dist[i] = math.MaxInt // ustalamy bardzo duże wartości
		parent[i] = -1
	}

	dist[source] = 0.0

	// Tworzymy kolejkę priorytetową i dodajemy do niej wierzchołek źródłowy.
	pq := make(priorityQueue, 1)
	pq[0] = &pqItem{
		node:     source,
		distance: 0.0,
		index:    0,
	}
	heap.Init(&pq)

	// Dla każdej relaksowanej krawędzi będziemy wyciągać z kolejki wierzchołek o najmniejszym dist.
	for pq.Len() > 0 {
		// Pobieramy wierzchołek o najmniejszej dotychczas znanej odległości
		current := heap.Pop(&pq).(*pqItem)
		u := current.node
		currentDist := current.distance

		// Jeśli odległość z kolejki jest większa niż w naszej tablicy dist[u],
		// to znaczy, że ten wierzchołek jest już "nieaktualny" i pomijamy go.
		if currentDist > dist[u] {
			continue
		}

		// Przeglądamy wszystkie potencjalne krawędzie wychodzące z u.
		for v := 0; v < g.Size; v++ {
			cost := g.data[u*g.Size+v]

			// Jeśli w macierzy jest 0 (i to oznacza brak krawędzi), to pomijamy.
			// UWAGA: jeśli 0 ma znaczyć "krawędź o koszcie 0", należy to dostosować.
			if cost < 0 {
				continue
			}

			alt := dist[u] + cost
			if alt < dist[v] {
				dist[v] = alt
				parent[v] = u
				// Dodajemy/aktualizujemy wierzchołek v w kolejce
				heap.Push(&pq, &pqItem{
					node:     v,
					distance: alt,
				})
			}
		}
	}

	return dist, parent, nil
}

// DijkstraDial implementuje algorytm Diala dla grafu o nieujemnych CAŁKOWITYCH wagach.
// Zakłada, że cost < 0 jest ignorowane (niedozwolone w teorii).
func DijkstraDial(g *Graph, source int) ([]int, []int, error) {
	n := g.Size
	if source < 0 || source >= n {
		return nil, nil, fmt.Errorf("invalid source: %d", source)
	}

	// 1. Wyznaczamy maksymalny koszt krawędzi
	maxC := MaxEdgeCost(g)
	if maxC == 0 {
		maxC = 1 // aby mieć co najmniej 1 kubełek, gdy wszystkie koszty=0
	}

	// W najprostszym (teoretycznym) wariancie Dial’s Algorithm
	// tworzymy buckets o rozmiarze (maxC * (n - 1) + 1).
	bucketCount := maxC*(n-1) + 1
	buckets := make([][]int, bucketCount)

	dist := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
		parent[i] = -1
	}
	dist[source] = 0

	// Wrzucamy wierzchołek źródłowy do kubełka o indeksie 0
	buckets[0] = append(buckets[0], source)

	// 2. Iteracja po kubełkach
	currentBucket := 0
	for currentBucket < bucketCount {
		for len(buckets[currentBucket]) > 0 {
			// Pobieramy ostatni wierzchołek z listy
			uIndex := len(buckets[currentBucket]) - 1
			u := buckets[currentBucket][uIndex]
			buckets[currentBucket] = buckets[currentBucket][:uIndex]

			// Jeżeli bieżąca odległość w dist[u] > currentBucket, pomijamy
			if dist[u] > currentBucket {
				continue
			}

			// Przeglądamy sąsiadów
			for v := 0; v < n; v++ {
				cost := g.data[u*n+v]
				if cost < 0 {
					// brak krawędzi (lub koszt ujemny - pomijamy)
					continue
				}
				newDist := dist[u] + cost
				if newDist < dist[v] {
					oldDist := dist[v]
					dist[v] = newDist
					parent[v] = u

					// Usuwamy v ze starego kubełka (jeśli nie był ∞)
					if oldDist < math.MaxInt {
						oldBucket := oldDist
						if oldBucket < bucketCount {
							// Usuwamy v z listy buckets[oldBucket]
							for iBucket, node := range buckets[oldBucket] {
								if node == v {
									buckets[oldBucket] = append(
										buckets[oldBucket][:iBucket],
										buckets[oldBucket][iBucket+1:]...,
									)
									break
								}
							}
						}
					}
					// Dodajemy v do nowego kubełka
					newBucket := newDist
					if newBucket < bucketCount {
						buckets[newBucket] = append(buckets[newBucket], v)
					}
				}
			}
		}
		currentBucket++
	}

	return dist, parent, nil
}

func DijkstraRadix(g *Graph, source int) ([]int, []int, error) {
	n := g.Size
	if source < 0 || source >= n {
		return nil, nil, fmt.Errorf("invalid source: %d", source)
	}

	dist := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt // reprezentacja "nieskończoności" w int
		parent[i] = -1
	}
	dist[source] = 0

	// Tworzymy RadixHeap i wstawiamy źródło
	rh := radix.NewRadixHeap()
	rh.Insert(source, 0)

	for !rh.Empty() {
		u, d := rh.PopMin()
		// Jeśli mamy już lepszą odległość, to wpis z kopca jest nieaktualny
		if d > dist[u] {
			continue
		}
		// Przechodzimy po sąsiadach
		for v := 0; v < n; v++ {
			cost := g.data[u*n+v]
			if cost < 0 {
				continue
			}
			newDist := d + cost
			if newDist < dist[v] {
				dist[v] = newDist
				parent[v] = u
				rh.Insert(v, newDist)
			}
		}
	}

	return dist, parent, nil
}
