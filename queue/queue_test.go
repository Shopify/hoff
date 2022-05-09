package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	q := NewQueue[int]()
	require.True(t, q.IsEmpty())

	for i := 0; i < 10_000; i++ {
		q.Enqueue(i)
	}
	require.False(t, q.IsEmpty())
	require.Equal(t, 10_000, q.Len())

	for i := 0; i < 10_000; i++ {
		dq, err := q.Dequeue()
		require.NoError(t, err)
		require.Equal(t, i, dq)
	}
	require.True(t, q.IsEmpty())

	_, err := q.Dequeue()
	require.Error(t, err)
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Peek()
	require.Error(t, err)

	q.Enqueue(1)
	v, err := q.Peek()
	require.NoError(t, err)
	require.Equal(t, 1, v)
}
