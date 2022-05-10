package utils

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"testing"
	"time"

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
	slice, err := generateRandomMapUint64(100_000, 100, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}

	for n := 0; n < b.N; n++ {
		Pluck(slice, generateIntInRange(0, 100), generateIntInRange(0, 100), generateIntInRange(0, 100))
	}
}

func generateRandomMapUint64(elements int, keys int, max int64) ([]map[int]uint64, error) {
	inputSlice := make([]map[int]uint64, elements)
	for i := 0; i < elements; i++ {
		m := make(map[int]uint64, keys)
		for k := 0; k < keys; k++ {
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

func generateIntInRange(min int, max int) int {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Intn(max-min+1) + min
}
