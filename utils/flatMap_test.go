package utils

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type flatMapTestCase struct {
	in  []string
	out []string
}

var testCases = []flatMapTestCase{
	{
		in:  []string{"abcd"},
		out: []string{"a", "b", "c", "d"},
	},
	{
		in:  []string{"abcd", "efg"},
		out: []string{"a", "b", "c", "d", "e", "f", "g"},
	},
}

type contextKey string

var key = contextKey("key")
var testKey = contextKey("T")

func TestFlatMap(t *testing.T) {
	for _, testCase := range testCases {
		out := FlatMap(testCase.in, splitString)
		require.Equal(t, testCase.out, out)
	}
}

func splitString(s string) []string {
	return strings.Split(s, "")
}

func ExampleFlatMap() {
	fmt.Println(FlatMap([]string{"abcd", "efg"}, func(s string) []int { return []int{len(s)} }))
	// Output: [4 3]
}

var benchmarkData = createStringSlice(100, 1000)

func BenchmarkFlatMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlatMap(benchmarkData, splitString)
	}
}

func createStringSlice(length, num int) []string {
	data := make([]string, num)
	value := strings.Repeat("a", length)

	for i := 0; i < num; i++ {
		data[i] = value
	}

	return data
}

func TestFlatMapContext(t *testing.T) {
	splitStringWithContext := func(c context.Context, s string) []string {
		require.Equal(t, "a_value", c.Value(key))

		return strings.Split(s, "")
	}
	ctx := context.WithValue(context.Background(), key, "a_value")

	for _, testCase := range testCases {
		out := FlatMapContext(ctx, testCase.in, splitStringWithContext)
		require.Equal(t, testCase.out, out)
		require.Equal(t, "a_value", ctx.Value(key))
	}
}
