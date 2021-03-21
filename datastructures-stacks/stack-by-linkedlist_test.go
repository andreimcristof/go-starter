package main

import "testing"

func TestNewStackMustHaveNilHead(t *testing.T) {
	s := NewLLStack()
	if s.head != nil {
		t.Error("head must exist and be nil")
	}
}

func TestNewStackInitializesSize(t *testing.T) {
	s := NewLLStack()

	if s.size != 0 {
		t.Error("expected size to be initialized to 0")
	}
}

func TestPushAndPeekLLStack(t *testing.T) {
	s := NewLLStack()

	s.push("a")
	s.push("b")
	s.push("c")

	actual := s.peek()

	if actual != "c" {
		t.Error("expected last item to be c")
	}
}

func TestPeekOnEmptyStack(t *testing.T) {
	s := NewLLStack()

	last := s.peek()

	if last != nil {
		t.Error("expected last to be nil in empty stack")
	}
}

func TestStackSize(t *testing.T) {
	s := NewLLStack()

	s.push("a")

	if s.len() != 1 {
		t.Error("expected length to be 1")
	}
}

func TestPopFromLLStackWithElements(t *testing.T) {
	s := NewLLStack()

	s.push("a")
	s.push("b")
	s.push("c")
	s.push("d")

	last := s.pop()
	if last != "d" {
		t.Error("expected last popped to be d")
	}

	if s.peek() != "c" {
		t.Error("expected peek to return c as last")
	}
}

func TestPopFromEmptyLLStack(t *testing.T) {
	s := NewLLStack()

	last := s.pop()

	if last != nil {
		t.Error("expected last from empty stack to be nil")
	}
}
