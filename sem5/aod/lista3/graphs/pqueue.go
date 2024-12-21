package graphs

import (
	"container/heap"
)

// pqItem przechowuje informację o wierzchołku (node) i aktualnej znanej odległości (distance).
type pqItem struct {
	node     int
	distance int
	index    int // index w kolejce (implementacja heap.Interface wymaga)
}

// priorityQueue to implementacja kolejki priorytetowej (min-heap).
type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push dodaje nowy element do kolejki.
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

// Pop usuwa element o najmniejszym priorytecie (najmniejsza distance).
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	// Zapobiegawczo usuwamy referencje, żeby nie trzymać w pamięci.
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update pozwala zaktualizować odległość do wierzchołka.
func (pq *priorityQueue) update(item *pqItem, distance int) {
	item.distance = distance
	heap.Fix(pq, item.index)
}
