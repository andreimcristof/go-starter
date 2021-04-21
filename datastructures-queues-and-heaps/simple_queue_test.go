package queue

import "testing"

func TestEnqueueAndDequeue(t *testing.T) {
	q := &Queue{}

	q.enqueue("a")
	q.enqueue("b")

	actual1 := q.dequeue()

	if actual1 != "a" {
		t.Error("expected first element removed to be a")
	}

	actual2 := q.dequeue()
	if actual2 != "b" {
		t.Error("expected second element removed to be b")
	}
}
