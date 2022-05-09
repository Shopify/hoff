package utils

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapConcurrent_simple(t *testing.T) {
	fn := func(ctx context.Context, val int, _ int) (string, error) {
		return strconv.Itoa(val), nil
	}

	results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

	require.Equal(t, 0, len(results.Errors()))
	require.Equal(t, []string{"1", "2", "3"}, results.Values())
}

func TestMapConcurrent_error(t *testing.T) {
	fn := func(ctx context.Context, val int, _ int) (string, error) {
		if val == 2 {
			return "", errors.New("I can't even")
		}
		return strconv.Itoa(val), nil
	}

	results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

	require.Equal(t, 1, len(results.Errors()))
	require.ErrorContains(t, results.Errors()[0], "I can't even")
	require.Equal(t, []string{"1", "3"}, results.Values())
}

func TestMapConcurrent_panic(t *testing.T) {
	fn := func(ctx context.Context, val int, _ int) (string, error) {
		if val == 2 {
			panic("I can't even")
		}
		return strconv.Itoa(val), nil
	}

	results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

	require.Equal(t, 1, len(results.Errors()))
	require.ErrorContains(t, results.Errors()[0], "MapConcurrent recovered from panic while processing value at index 1: I can't even") //nolint:lll
	require.Equal(t, []string{"1", "3"}, results.Values())
}
