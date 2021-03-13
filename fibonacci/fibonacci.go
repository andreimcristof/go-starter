package main

func fibOne(position int) int {
	seq := []int{0, 1}

	for len(seq) <= position {
		last := seq[len(seq)-1]
		beforeLast := seq[len(seq)-2]
		seq = append(seq, last+beforeLast)
	}

	return seq[len(seq)-1]
}

func fibTwo(position int, seq []int) int {
	if seq == nil {
		seq = []int{0, 1}
	}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
		return fibTwo(position, seq)
	}

	return seq[len(seq)-1]
}

func fibThree(n int) int {
	if n < 2 {
		return n
	}

	a, b := fibThree(n-1), fibThree(n-2)
	return a + b
}
