package data

import "log"

type list struct {
	item int
	next *list
}

// Queue is int queue
type Queue struct {
	first *list
	last  *list
}

// NewQueue create a new queue
func NewQueue() *Queue {
    return &Queue{first: nil, last: nil}
}

// Enqueue add new element to the queue
func (q *Queue) Enqueue(e int) {
	oldlast := q.last
	q.last = &list{item: e, next: nil}
	if q.first == nil {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
}

// Dequeue returns and removes least recent queue element
func (q *Queue) Dequeue() int {
	if q.first == nil {
		log.Fatal("queue is empty")
	}
	oldfirst := q.first
	q.first = oldfirst.next
	if q.first == nil {
		q.last = nil
	}
	return oldfirst.item
}

func (q *Queue) IsEmpty() bool {
    return q.first == nil
}
