package generics

import (
	"errors"
)

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{values: []T{}}
}

func (s *Stack[any]) Push(item any) {
	s.values = append(s.values, item)
}

func (s *Stack[any]) Pop() (any, error) {
	var empty any
	length := len(s.values)
	if length < 1 {
		return empty, errors.New("empty stack")
	}
	element := s.values[length-1]
	s.values = s.values[:length-1]
	return element, nil
}
