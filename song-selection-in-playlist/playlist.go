package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SongList type
type SongList []*Song

// random return a random value from the list
func (s SongList) random() *Song {
	if len(s) == 0 {
		return nil
	}

	// this seed guarantees a new value every time it's called
	rand.Seed(time.Now().UnixNano())
	rndPos := rand.Intn(len(s))
	return s[rndPos]
}

// Playlist type
type Playlist struct {
	songs      SongList
	nowPlaying *Song
}

// NewPlaylist create a new playlist - simulated constructor
func NewPlaylist(songs []Song) *Playlist {
	p := Playlist{}

	for i := range songs {
		p.songs = append(p.songs, &songs[i])
	}

	return &p
}

func (p *Playlist) sort() {
	sort.Sort(SortByLikes(*p))
}

func (p *Playlist) print() {
	fmt.Println("\nYour Playlist:")

	for _, song := range p.songs {
		song.print()
	}
}

func (p *Playlist) getNextToPlay() *Song {
	var otherCategorySongs SongList
	if p.nowPlaying != nil {
		for _, s := range p.songs {
			if s.title != p.nowPlaying.title && s.likes < 2 {
				if s.category == p.nowPlaying.category {
					return s
				}
				otherCategorySongs = append(otherCategorySongs, s)
			}
		}
	}

	// if current category was listened to twice, try another category
	if len(otherCategorySongs) > 0 {
		return otherCategorySongs.random()
	}

	// all was listened to twice, or nothing was running in the playlist. return a random song
	p.resetLikes()
	return p.songs.random()
}

func (p *Playlist) resetLikes() {
	for _, s := range p.songs {
		s.likes = 0
	}
}

func (p *Playlist) play(s *Song) {
	p.nowPlaying = s
	fmt.Println(fmt.Sprintf("\nNow playing: %s (%s)\npress 1 / 0 to like / dislike", s.title, s.category.string()))
}
