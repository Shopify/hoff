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

func TestFlatMap(t *testing.T) {
	for _, testCase := range testCases {
		out := FlatMap(testCase.in, splitString)
		require.Equal(t, testCase.out, out)
	}
}

func ExampleFlatMap() {
	fmt.Println(FlatMap([]string{"abcd", "efg"}, func(s string) []int { return []int{len(s)} }))
	// Output: [4 3]
}

var benchmarkData = createStringSlice(100, 1000)
var benchmarkResult []string
var benchmarkError error

func BenchmarkFlatMap(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = FlatMap(benchmarkData, splitString)
	}
	benchmarkResult = r
}

func TestFlatMapErrorSucceeds(t *testing.T) {
	out, err := FlatMapError(testCases[0].in, splitStringWithError)

	require.NoError(t, err)
	require.Equal(t, testCases[0].out, out)
}

func TestFlatMapErrorFails(t *testing.T) {
	out, err := FlatMapError(testCases[1].in, splitStringWithError)

	require.Error(t, err)
	require.Len(t, out, 0)
}

func ExampleFlatMapError() {
	properPrefixes := func(s string) ([]string, error) {
		if len(s) < 2 {
			return nil, fmt.Errorf("String '%s' has no proper prefixes", s)
		}

		res := make([]string, len(s)-1)
		for i := 1; i < len(s); i++ {
			res[i-1] = s[0:i]
		}

		return res, nil
	}

	var err error
	prefixes, _ := FlatMapError([]string{"abcd", "efg"}, properPrefixes)
	fmt.Println(prefixes)
	prefixes, err = FlatMapError([]string{"abcd", "x", "efg"}, properPrefixes)
	fmt.Println(err)
	// Output:
	// [a ab abc e ef]
	// String 'x' has no proper prefixes
}

func BenchmarkFlatMapError(b *testing.B) {
	var r []string
	var err error
	for i := 0; i < b.N; i++ {
		r, err = FlatMapError(benchmarkData, splitStringWithError)
	}
	benchmarkResult = r
	benchmarkError = err
}

func TestFlatMapContext(t *testing.T) {
	fn := func(c context.Context, s string) []string {
		require.Equal(t, "a_value", c.Value(key))

		return splitString(s)
	}
	ctx := context.WithValue(context.Background(), key, "a_value")

	for _, testCase := range testCases {
		out := FlatMapContext(ctx, testCase.in, fn)
		require.Equal(t, testCase.out, out)
		require.Equal(t, "a_value", ctx.Value(key))
	}
}

func TestFlatMapContextErrorSucceeds(t *testing.T) {
	fn := func(c context.Context, s string) ([]string, error) {
		require.Equal(t, "a_value", c.Value(key))

		return splitStringWithContextAndError(c, s)
	}
	ctx := context.WithValue(context.Background(), key, "a_value")

	out, err := FlatMapContextError(ctx, testCases[0].in, fn)
	require.NoError(t, err)
	require.Equal(t, testCases[0].out, out)
	require.Equal(t, "a_value", ctx.Value(key))
}

func TestFlatMapContextErrorFails(t *testing.T) {
	fn := func(c context.Context, s string) ([]string, error) {
		require.Equal(t, "a_value", c.Value(key))

		return splitStringWithContextAndError(c, s)
	}
	ctx := context.WithValue(context.Background(), key, "a_value")

	out, err := FlatMapContextError(ctx, testCases[1].in, fn)
	require.Error(t, err)
	require.Len(t, out, 0)
	require.Equal(t, "a_value", ctx.Value(key))
}

func splitString(s string) []string {
	return strings.Split(s, "")
}

func splitStringWithError(s string) ([]string, error) {
	if len(s) < 4 {
		return nil, fmt.Errorf("String '%s' too short", s)
	}
	return strings.Split(s, ""), nil
}

func splitStringWithContextAndError(ctx context.Context, s string) ([]string, error) {
	if len(s) < 4 {
		return nil, fmt.Errorf("String '%s' too short", s)
	}
	return strings.Split(s, ""), nil
}

func createStringSlice(length, num int) []string {
	data := make([]string, num)
	value := strings.Repeat("a", length)

	for i := 0; i < num; i++ {
		data[i] = value
	}

	return data
}
