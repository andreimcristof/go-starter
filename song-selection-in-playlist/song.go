package main

import "fmt"

// Category enum type
type Category int

const (
	pop Category = iota
	jazz
	rock
)

func (c Category) string() string {
	switch c {
	case jazz:
		return "Jazz"
	case pop:
		return "Pop"
	case rock:
		return "Rock"
	default:
		return ""
	}
}

// Song type
type Song struct {
	title    string
	category Category
	likes    int
}

func (s *Song) print() {
	msg := fmt.Sprintf("%s (%s), likes: %d,", s.title, s.category.string(), s.likes)
	fmt.Println(msg)
}

func (s *Song) setLiked(liked bool) {
	if liked {
		s.likes++
	} else {
		if s.likes > 0 {
			s.likes--
		}
	}
}
