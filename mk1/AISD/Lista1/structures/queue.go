package structures

import "errors"

type Queue[T any] struct {
	memory []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.memory = append(q.memory, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.memory) == 0 {
		var zero T
		return zero, errors.New("empty queue")
	}
	ret := q.memory[0]
	q.memory = q.memory[1:]
	return ret, nil
}
