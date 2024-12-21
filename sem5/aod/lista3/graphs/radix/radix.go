package radix

import "math/bits"

// item przechowuje parę: wierzchołek (lub dowolny identyfikator) i priorytet (distance).
type item struct {
	node int
	dist int
}

// RadixHeap to nasz „koszyk” warstwowy.
type RadixHeap struct {
	buckets     [][]item // tablica kubełków
	lastDeleted int      // priorytet ostatnio usuniętego (PopMin)
	size        int      // liczba elementów
}

// NewRadixHeap tworzy pusty RadixHeap.
// Można dać np. 33 kubełki dla liczb 32-bit (indeks 0..32).
func NewRadixHeap() *RadixHeap {
	// Zamiast 33 możliwe jest 32, ale często dodaje się jeden „zapasowy” kubełek
	const bucketCount = 33
	return &RadixHeap{
		buckets: make([][]item, bucketCount),
		// lastDeleted = 0 na starcie (zakładamy brak elementów)
		lastDeleted: 0,
		size:        0,
	}
}

// bucketIndex oblicza numer kubełka, do którego wpadnie wartość `dist`
// zgodnie z klasycznym wzorem: jeśli dist == lastDeleted => i=0,
// wpp. i = floor(log2(dist ^ lastDeleted)) + 1
func (rh *RadixHeap) bucketIndex(dist int) int {
	if dist == rh.lastDeleted {
		return 0
	}
	// x := dist ^ lastDeleted
	x := dist ^ rh.lastDeleted
	// pozycja najstarszego ustawionego bitu (log2) => bits.Len32(x) - 1
	// +1 bo bucket[0] rezerwujemy na równych lastDeleted
	return bits.Len(uint(x))
}

// Insert wstawia nowy (node, dist) do RadixHeap.
// Założenie: dist >= lastDeleted (bo nie możemy mieć klucza mniejszego
// niż już usunięty – w kontekście Dijkstry to i tak się nie zdarza,
// bo dystanse relaksowane zawsze rosną lub się poprawiają z góry,
// ale muszą być >= lastDeleted).
func (rh *RadixHeap) Insert(n, dist int) {
	if dist < rh.lastDeleted {
		// W teorii można tu zrobić panic lub po cichu zignorować,
		// ale w poprawnym algorytmie Dijkstry to się nie zdarzy.
		dist = rh.lastDeleted
	}
	bi := rh.bucketIndex(dist)
	rh.buckets[bi] = append(rh.buckets[bi], item{node: n, dist: dist})
	rh.size++
}

// Empty mówi, czy RadixHeap jest pusty
func (rh *RadixHeap) Empty() bool {
	return rh.size == 0
}

// PopMin zwraca parę (node, dist) o najmniejszym dist.
// Zasada: najpierw sprawdzamy bucket[0]. Jeśli pusty –
// znajdujemy pierwszy niepusty i robimy tzw. "redistribucję".
func (rh *RadixHeap) PopMin() (int, int) {
	if rh.size == 0 {
		// Zwracamy np. (-1, -1) albo panikujemy
		return -1, -1
	}

	// Jeśli bucket[0] jest niepusty, tam na pewno jest minimum (dist == lastDeleted)
	if len(rh.buckets[0]) > 0 {
		u := rh.buckets[0][len(rh.buckets[0])-1]             // weźmy ostatni
		rh.buckets[0] = rh.buckets[0][:len(rh.buckets[0])-1] // usuń go z kubełka
		rh.size--
		return u.node, u.dist
	}

	// bucket[0] jest pusty, więc musimy znaleźć pierwszy niepusty kubełek > 0
	var idx int
	for i := 1; i < len(rh.buckets); i++ {
		if len(rh.buckets[i]) > 0 {
			idx = i
			break
		}
	}

	// Znajdujemy minimalną wartość dist w tym kubełku
	minDist := rh.buckets[idx][0].dist
	for _, it := range rh.buckets[idx] {
		if it.dist < minDist {
			minDist = it.dist
		}
	}

	// Ustawiamy lastDeleted = minDist
	rh.lastDeleted = minDist

	// Teraz przenosimy wszystkie elementy z bucket[idx]
	// do właściwych kubełków (bo nowy lastDeleted się zmienił).
	oldItems := rh.buckets[idx]
	rh.buckets[idx] = rh.buckets[idx][:0] // czyścimy kubełek
	for _, it := range oldItems {
		bi := rh.bucketIndex(it.dist)
		rh.buckets[bi] = append(rh.buckets[bi], it)
	}

	// Po redistribucji w bucket[0] *powinny* być teraz wszystkie elementy,
	// które mają dist == lastDeleted (czyli równe minDist).
	// Możemy więc z bucket[0] wyjąć minDist
	u := rh.buckets[0][len(rh.buckets[0])-1] // weźmy ostatni
	rh.buckets[0] = rh.buckets[0][:len(rh.buckets[0])-1]
	rh.size--

	return u.node, u.dist
}
