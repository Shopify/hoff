package queue

import "fmt"

type boundedQueue[T any] struct {
	capacity int
	queue    *Queue[T]
}

func NewBoundedQueue[T any](capacity int) *boundedQueue[T] {
	bq := &boundedQueue[T]{
		capacity: capacity,
		queue:    NewQueue[T](),
	}
	return bq
}

func (bq *boundedQueue[T]) IsFull() bool {
	return bq.Len() == bq.Cap()
}

func (bq *boundedQueue[T]) IsEmpty() bool {
	return bq.queue.IsEmpty()
}

func (bq *boundedQueue[T]) Cap() int {
	return (*bq).capacity
}

func (bq *boundedQueue[T]) Len() int {
	return bq.queue.Len()
}

func (bq *boundedQueue[T]) Enqueue(values ...T) error {
	if len(values)+bq.queue.Len() > (*bq).capacity {
		return fmt.Errorf("queue would overflow")
	}
	(*bq.queue).Enqueue(values...)
	return nil
}

func (bq *boundedQueue[T]) Dequeue() (T, error) {
	return (*bq.queue).Dequeue()
}

func (bq *boundedQueue[T]) Peek() (T, error) {
	return (*bq.queue).Peek()
}
