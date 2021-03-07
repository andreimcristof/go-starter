package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	selection := []Song{
		{
			title:    "Metallica - One",
			likes:    0,
			category: rock,
		},
		{
			title:    "Lenny Kravitz - Fly Away",
			likes:    0,
			category: rock,
		},
		{
			title:    "Chillout Lounge",
			likes:    0,
			category: jazz,
		},
		{
			title:    "Miles Davis - The Doo-Bop Song",
			likes:    0,
			category: jazz,
		},
		{
			title:    "Michael Jackson - Billie Jean",
			likes:    0,
			category: pop,
		},
		{
			title:    "Madonna - Frozen",
			likes:    0,
			category: pop,
		},
	}

	p := NewPlaylist(selection)

	p.print()

	// reads from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// we use the time.NewTicker to simulate a running playlist
	tck := time.NewTicker(50 * time.Millisecond)

	for range tck.C {
		next := p.getNextToPlay()
		p.play(next)

		for scanner.Scan() {
			if scanner.Text() == "x" {
				fmt.Println("exiting player...")
				os.Exit(0)
			}

			option, _ := strconv.ParseInt(scanner.Text(), 0, 64)
			next.setLiked(option == 1)

			// break the for loop for user input, so that the playlist (outer for loop) can continue
			break
		}

		p.sort()
		p.print()
	}
}
