package main

// Stack type based on a slice
type Stack struct {
	items []interface{}
}

// NewStack ctor for slice-based Stack
func NewStack() *Stack {
	s := Stack{}
	s.items = make([]interface{}, 0)
	return &s
}

func (s *Stack) push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	last := s.items[len(s.items)-1]
	return last
}
