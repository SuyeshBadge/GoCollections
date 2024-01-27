// Package queue provides a basic queue data structure implementation.
package queue

import (
	"errors"
)

// Queue represents a simple queue structure.
type Queue struct {
	values []interface{}
	size   int
}

// NewQueue creates and returns a new empty queue.
func NewQueue() *Queue {
	return &Queue{
		size: 0,
	}
}

// Enqueue adds a new element to the end of the queue.
func (q *Queue) Enqueue(value interface{}) {
	q.values = append(q.values, value)
	q.size++
}

// Dequeue removes and returns the element from the front of the queue.
// It returns an error if the queue is empty.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}

	lastElement := q.values[0]
	q.values = q.values[1:]
	q.size--

	return lastElement, nil
}

// Head returns the element at the front of the queue without removing it.
// It returns an error if the queue is empty.
func (q *Queue) Head() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.values[0], nil
}

// Tail returns the element at the end of the queue without removing it.
// It returns an error if the queue is empty.
func (q *Queue) Tail() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.values[q.size-1], nil
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return q.size
}

// Clear removes all elements from the queue, making it empty.
func (q *Queue) Clear() {
	q.values = nil
	q.size = 0
}
