package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResults(t *testing.T) {
	ok := Results[int]{
		{1, nil},
		{2, nil},
		{3, nil},
	}
	err := Results[int]{
		{1, nil},
		{2, nil},
		{0, errors.New("fizz")},
	}

	t.Run("HasErrors", func(t *testing.T) {
		require.False(t, ok.HasError())
		require.True(t, err.HasError())
	})

	t.Run("Errors", func(t *testing.T) {
		require.Equal(t, []error{nil, nil, nil}, ok.Errors())
		require.Equal(t, []error{nil, nil, errors.New("fizz")}, err.Errors())
	})

	t.Run("Values", func(t *testing.T) {
		require.Equal(t, []int{1, 2, 3}, ok.Values())
		require.Equal(t, []int{1, 2, 0}, err.Values())
	})
}
