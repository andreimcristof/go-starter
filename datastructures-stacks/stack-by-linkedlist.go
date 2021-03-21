package main

// Item type for LLStack
type Item struct {
	data interface{}
	next *Item
}

// LLStack - stack implemented with linked list
type LLStack struct {
	head *Item
	size int
}

// NewLLStack ctor for Linked List-based Stack
func NewLLStack() *LLStack {
	s := LLStack{}
	s.size = 0
	return &s
}

// push - the head of linked list is always the top element
func (s *LLStack) push(data interface{}) {
	item := &Item{data: data, next: s.head}
	s.head = item
	s.size++
}

func (s *LLStack) pop() interface{} {
	if s.head == nil {
		return nil
	}
	h := s.head
	s.head = s.head.next
	return h.data
}

func (s *LLStack) peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.data
}

func (s *LLStack) len() int {
	return s.size
}
