package internal

type Stack[T any] struct {
	items []*T
}

func (s *Stack[T]) New(initial uint) {
	s.items = make([]*T, initial)

}

func (s *Stack[T]) Push(item *T) {
	s.items = append(s.items, item)

}

func (s *Stack[T]) Pop() *T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}
