package utils

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPluckStringBasic(t *testing.T) {
	output := Pluck([]map[string]any{
		{"foo": 1, "bar": 2, "shop": 3},
		{"foo": 4, "bar": 5, "shop": 6},
	}, "foo", "bar")
	require.Equal(t, [][]any{{1, 2}, {4, 5}}, output, "2 keys extracted")
}

func TestPluckStringNil(t *testing.T) {
	output := Pluck([]map[string]any{
		{"foo": 1, "bar": 2, "shop": 3},
		{"foo": 4, "bar": 5},
	}, "foo", "shop")
	require.Equal(t, [][]any{{1, 3}, {4, nil}}, output, "2 keys extracted with nil value")
}

func TestPluckStringSingleKey(t *testing.T) {
	output := Pluck([]map[string]any{
		{"foo": 1, "bar": 2, "shop": 3},
		{"foo": 4, "bar": 5},
	}, "foo")
	require.Equal(t, [][]any{{1}, {4}}, output, "1 keys extracted")
}

func BenchmarkPluck(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slice, err := generateRandomMapUint64(10_000, 1000, 1_000_000)
		if err != nil {
			b.Errorf("Error in generating random slice")
		}
		Pluck(slice, 1, 2, 3)
	}
}

func generateRandomMapUint64(elements int, keys int, max int64) ([]map[int]uint64, error) {
	inputSlice := make([]map[int]uint64, elements)
	for i := 0; i < elements; i++ {
		m := make(map[int]uint64, keys)
		for k := 0; k < keys; i++ {
			value, err := rand.Int(rand.Reader, big.NewInt(max))
			if err != nil {
				return nil, err
			}
			m[k] = value.Uint64()
		}
		inputSlice[i] = m
	}
	return inputSlice, nil
}
