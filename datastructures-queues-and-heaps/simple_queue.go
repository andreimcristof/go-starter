package queue

type Queue struct {
	items []*interface{}
}

// enqueue add to the back of the queue
func (q *Queue) enqueue(item interface{}) {
	q.items = append(q.items, &item)
}

func (q *Queue) dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	// we store a reference to the first item
	first := q.items[0]

	// we reassign the slice of items to remove first
	q.items = q.items[1:]

	return *first
}
