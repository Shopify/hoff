package queue

import "fmt"

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	queue := make(Queue[T], 0)
	return &queue
}

func (q *Queue[T]) Enqueue(values ...T) {
	for _, i := range values {
		*q = append(*q, i)
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Len() == 0 {
		var t T
		return t, fmt.Errorf("queue is empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Len() == 0 {
		var t T
		return t, fmt.Errorf("queue is empty")
	}
	return (*q)[0], nil
}
