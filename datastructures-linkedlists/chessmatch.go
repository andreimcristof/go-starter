package main

import "fmt"

// Move the element of list
type Move struct {
	data string
	next *Move
}

// NewMove ctor for Move
func NewMove(data string, next *Move) *Move {
	return &Move{data: data, next: next}
}

// ChessMatch the linked list storing chess moves
type ChessMatch struct {
	head *Move
}

// NewChessMatch ctor for chess game
func NewChessMatch() *ChessMatch {
	return &ChessMatch{}
}

func (c *ChessMatch) getAt(index int) *Move {
	crt := c.head
	pos := 0

	for pos < index && crt != nil {
		crt = crt.next
		pos++
	}

	return crt
}

func (c *ChessMatch) getLast() *Move {
	if c.head == nil {
		return nil
	}

	current := c.head

	for current != nil && current.next != nil {
		current = current.next
	}
	return current
}

func (c *ChessMatch) insertAt(index int, data string) {
	if c.head == nil {
		c.head = NewMove(data, nil)
		return
	}

	if index == 0 {
		h := c.head
		c.head = NewMove(data, h)
		return
	}

	prev := c.getAt(index - 1)
	if prev == nil {
		prev = c.getLast()
	}
	prev.next = NewMove(data, prev.next)
}

func (c *ChessMatch) removeAt(index int) {
	if c.head == nil {
		return
	}

	if index == 0 {
		c.head = c.head.next
		return
	}

	prev := c.getAt(index - 1)
	if prev != nil && prev.next != nil {
		prev.next = prev.next.next
	}

}

func (c *ChessMatch) forEach(predicateFn func(*Move)) {
	if c.head == nil {
		return
	}

	current := c.head
	for current != nil {
		predicateFn(current)
		current = current.next
	}
}

func (c *ChessMatch) replayMatch() {
	printMove := func(t *Move) { fmt.Println(t.data) }
	c.forEach(printMove)
}
