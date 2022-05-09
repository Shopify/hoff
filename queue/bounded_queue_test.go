package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoundedQueueEnqueue(t *testing.T) {
	bq := NewBoundedQueue[int](3)
	require.Equal(t, 3, bq.Cap())

	err := bq.Enqueue(1, 2, 3)
	require.NoError(t, err)
	require.Equal(t, 3, bq.Len())

	err = bq.Enqueue(4)
	require.Error(t, err)
	require.Equal(t, 3, bq.Len())
}

func TestBoundedQueueDequeue(t *testing.T) {
	bq := NewBoundedQueue[int](3)
	require.Equal(t, 3, bq.Cap())

	err := bq.Enqueue(1, 2, 3)
	require.NoError(t, err)
	require.Equal(t, 3, bq.Len())
	require.True(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err := bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 1, v)
	require.Equal(t, bq.Len(), 2)
	require.False(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 2, v)
	require.Equal(t, bq.Len(), 1)
	require.False(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 3, v)
	require.Equal(t, bq.Len(), 0)
	require.False(t, bq.IsFull())
	require.True(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.Error(t, err)
}
