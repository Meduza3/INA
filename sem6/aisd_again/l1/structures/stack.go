package structures

import "errors"

type Struct[T any] struct {
	memory []T
}

func (s *Struct[T]) Push(val T) {
	s.memory = append(s.memory, val)
}

func (s *Struct[T]) Pop() (T, error) {
	if len(s.memory) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	val := s.memory[len(s.memory)-1]
	s.memory = s.memory[:len(s.memory)-1]
	return val, nil
}
