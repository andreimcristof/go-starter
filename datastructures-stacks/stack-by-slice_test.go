package main

import "testing"

func TestEnsureStackInitWithEmptySlice(t *testing.T) {
	s := NewStack()
	if s.items == nil {
		t.Error("expected stack to be initialized with empty items slice")
	}
}

func TestPushAndPeek(t *testing.T) {
	s := NewStack()

	s.push("a")
	s.push("b")
	s.push("c")

	actualTopItem := s.peek()

	if actualTopItem != "c" {
		t.Error("expected c to be top of stack")
	}
}

func TestPop(t *testing.T) {
	s := NewStack()

	s.push("a")
	s.push("b")
	s.push("c")
	s.push("d")

	last := s.pop()
	if last != "d" {
		t.Error("expected last to be d")
	}

	if s.peek() != "c" {
		t.Error("expected peek to return c")
	}
}

func TestPeekAtEmptyStack(t *testing.T) {
	s := NewStack()

	last := s.peek()

	if last != nil {
		t.Error("expected last item in empty stack to be nil")
	}
}

func TestPopWithEmptyStack(t *testing.T) {
	s := NewStack()

	last := s.pop()

	if last != nil {
		t.Error("expected las to be nil")
	}
}
