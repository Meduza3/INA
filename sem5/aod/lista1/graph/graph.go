package graph

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Graph represents a graph using adjacency lists
type Graph struct {
	Directed bool
	N        int
	M        int
	Adj      [][]int
}

// ReadGraph wczytuje graf z pliku
func ReadGraph(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Odczyt flagi: D lub U
	if !scanner.Scan() {
		return nil, fmt.Errorf("nieoczekiwany koniec pliku podczas odczytu typu grafu")
	}
	flagLine := strings.TrimSpace(scanner.Text())
	var directed bool
	if flagLine == "D" {
		directed = true
	} else if flagLine == "U" {
		directed = false
	} else {
		return nil, fmt.Errorf("nieprawidłowa flaga typu grafu: %s", flagLine)
	}

	// Odczyt liczby wierzchołków
	if !scanner.Scan() {
		return nil, fmt.Errorf("nieoczekiwany koniec pliku podczas odczytu liczby wierzchołków")
	}
	nStr := strings.TrimSpace(scanner.Text())
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return nil, fmt.Errorf("nieprawidłowa liczba wierzchołków: %s", nStr)
	}

	// Odczyt liczby krawędzi
	if !scanner.Scan() {
		return nil, fmt.Errorf("nieoczekiwany koniec pliku podczas odczytu liczby krawędzi")
	}
	mStr := strings.TrimSpace(scanner.Text())
	m, err := strconv.Atoi(mStr)
	if err != nil {
		return nil, fmt.Errorf("nieprawidłowa liczba krawędzi: %s", mStr)
	}

	// Inicjalizacja listy sąsiedztwa
	adj := make([][]int, n+1) // 1-based indexing

	// Odczyt m krawędzi
	for i := 0; i < m; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("nieoczekiwany koniec pliku podczas odczytu krawędzi")
		}
		edgeLine := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(edgeLine)
		if len(parts) != 2 {
			return nil, fmt.Errorf("nieprawidłowa definicja krawędzi: %s", edgeLine)
		}
		u, err1 := strconv.Atoi(parts[0])
		v, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("nieprawidłowy wierzchołek w krawędzi: %s", edgeLine)
		}
		if u < 1 || u > n || v < 1 || v > n {
			return nil, fmt.Errorf("wierzchołek poza zakresem w krawędzi: %s", edgeLine)
		}
		adj[u] = append(adj[u], v)
		if !directed {
			adj[v] = append(adj[v], u)
		}
	}

	return &Graph{
		Directed: directed,
		N:        n,
		M:        m,
		Adj:      adj,
	}, nil
}

// BFS wykonuje przeszukiwanie wszerz zaczynając od wierzchołka start
func (g *Graph) BFS(start int) ([]int, []int) {
	visited := make([]bool, g.N+1)
	order := []int{}
	parent := make([]int, g.N+1)

	queue := list.New()
	queue.PushBack(start)
	visited[start] = true
	parent[start] = -1 // korzeń nie ma rodzica

	for queue.Len() > 0 {
		element := queue.Front()
		u := element.Value.(int)
		queue.Remove(element)
		order = append(order, u)

		for _, v := range g.Adj[u] {
			if !visited[v] {
				visited[v] = true
				parent[v] = u
				queue.PushBack(v)
			}
		}
	}

	return order, parent
}

// DFS wykonuje przeszukiwanie wgłąb zaczynając od wierzchołka start
func (g *Graph) DFS(start int) ([]int, []int) {
	visited := make([]bool, g.N+1)
	order := []int{}
	parent := make([]int, g.N+1)

	stack := list.New()
	stack.PushBack(start)
	parent[start] = -1

	for stack.Len() > 0 {
		element := stack.Back()
		u := element.Value.(int)
		stack.Remove(element)

		if !visited[u] {
			visited[u] = true
			order = append(order, u)

			// Dodaj sąsiadów w odwrotnej kolejności, aby zachować kolejność przeszukiwania
			for i := len(g.Adj[u]) - 1; i >= 0; i-- {
				v := g.Adj[u][i]
				if !visited[v] {
					stack.PushBack(v)
					if parent[v] == 0 { // ustaw rodzica tylko jeśli nie jest ustawiony
						parent[v] = u
					}
				}
			}
		}
	}

	return order, parent
}

// SCCResult przechowuje wynik algorytmu SCC
type SCCResult struct {
	Components [][]int
}

// SCC wykonuje algorytm Kosaraju na grafie i zwraca silnie spójne składowe
func (g *Graph) SCC() SCCResult {
	// Krok 1: Wykonaj DFS i przechowaj wierzchołki według czasu zakończenia
	visited := make([]bool, g.N+1)
	var finishOrder []int

	for u := 1; u <= g.N; u++ {
		if !visited[u] {
			g.dfsForSCC(u, visited, &finishOrder)
		}
	}

	// Krok 2: Odwróć graf
	revGraph := g.Reverse()

	// Krok 3: Wykonaj DFS na odwróconym grafie w kolejności malejącego czasu zakończenia
	visited = make([]bool, g.N+1)
	var components [][]int

	for i := len(finishOrder) - 1; i >= 0; i-- {
		u := finishOrder[i]
		if !visited[u] {
			var component []int
			revGraph.dfsCollect(u, visited, &component)
			components = append(components, component)
		}
	}

	return SCCResult{Components: components}
}

// dfsForSCC jest pomocniczą funkcją DFS używaną w algorytmie Kosaraju
func (g *Graph) dfsForSCC(u int, visited []bool, finishOrder *[]int) {
	visited[u] = true
	for _, v := range g.Adj[u] {
		if !visited[v] {
			g.dfsForSCC(v, visited, finishOrder)
		}
	}
	*finishOrder = append(*finishOrder, u)
}

// Reverse zwraca odwrócony graf
func (g *Graph) Reverse() *Graph {
	revAdj := make([][]int, g.N+1)
	for u := 1; u <= g.N; u++ {
		for _, v := range g.Adj[u] {
			revAdj[v] = append(revAdj[v], u)
		}
	}
	return &Graph{
		Directed: g.Directed,
		N:        g.N,
		M:        g.M,
		Adj:      revAdj,
	}
}

// dfsCollect jest pomocniczą funkcją DFS używaną do zbierania wierzchołków składowej
func (g *Graph) dfsCollect(u int, visited []bool, component *[]int) {
	visited[u] = true
	*component = append(*component, u)
	for _, v := range g.Adj[u] {
		if !visited[v] {
			g.dfsCollect(v, visited, component)
		}
	}
}

// IsBipartiteResult przechowuje wynik sprawdzania dwudzielności
type IsBipartiteResult struct {
	IsBipartite bool
	Partition   [2][]int
}

// IsBipartite sprawdza, czy graf jest dwudzielny. Jeśli tak, zwraca podział na dwa zbiory.
func (g *Graph) IsBipartite() IsBipartiteResult {
	// Inicjalizacja kolorów: -1 oznacza nieodwiedzony
	colors := make([]int, g.N+1)
	for i := range colors {
		colors[i] = -1
	}

	isBipartite := true
	var partition [2][]int

	// Funkcja BFS dla jednego komponentu
	bfs := func(start int) bool {
		queue := list.New()
		queue.PushBack(start)
		colors[start] = 0
		partition[0] = append(partition[0], start)

		for queue.Len() > 0 {
			element := queue.Front()
			u := element.Value.(int)
			queue.Remove(element)

			for _, v := range g.Adj[u] {
				if colors[v] == -1 {
					colors[v] = 1 - colors[u]
					partition[colors[v]] = append(partition[colors[v]], v)
					queue.PushBack(v)
				} else if colors[v] == colors[u] {
					return false
				}
			}
		}
		return true
	}

	// Sprawdzenie każdego wierzchołka (dla nieskierowanych grafów, aby uwzględnić wszystkie komponenty)
	for u := 1; u <= g.N; u++ {
		if colors[u] == -1 {
			if !bfs(u) {
				isBipartite = false
				break
			}
		}
	}

	if isBipartite {
		return IsBipartiteResult{
			IsBipartite: true,
			Partition:   partition,
		}
	}
	return IsBipartiteResult{
		IsBipartite: false,
	}
}
