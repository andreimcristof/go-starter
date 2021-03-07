package main

// SortByLikes sorter for playlist songs - by number of likes
type SortByLikes Playlist

// Len length of array
func (sorter SortByLikes) Len() int {
	return len(sorter.songs)
}

// Swap swap the elements behavior
func (sorter SortByLikes) Swap(a, b int) {
	sorter.songs[a], sorter.songs[b] = sorter.songs[b], sorter.songs[a]
}

// Less comparing two elements in the array
func (sorter SortByLikes) Less(a, b int) bool {
	return sorter.songs[a].likes > sorter.songs[b].likes
}
