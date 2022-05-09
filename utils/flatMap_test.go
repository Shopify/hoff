package utils

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type FlatMapTestCase struct {
	In    []string
	Out   []string
	Fn    func(s string) []string
	FnCtx func(ctx context.Context, s string) []string
}

var testCases = []FlatMapTestCase{
	{
		In:    []string{"abcd"},
		Out:   []string{"a", "b", "c", "d"},
		Fn:    splitString,
		FnCtx: splitStringWithContext,
	},
	{
		In:    []string{"abcd", "efg"},
		Out:   []string{"a", "b", "c", "d", "e", "f", "g"},
		Fn:    splitString,
		FnCtx: splitStringWithContext,
	},
}

type contextKey string

var key = contextKey("key")
var testKey = contextKey("T")

func TestFlatMap(t *testing.T) {
	for _, testCase := range testCases {
		out := FlatMap(testCase.In, testCase.Fn)
		require.Equal(t, testCase.Out, out)
	}
}

func TestFlatMapContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), key, "value")
	ctx = context.WithValue(ctx, testKey, t)

	for _, testCase := range testCases {
		out := FlatMapContext(ctx, testCase.In, testCase.FnCtx)
		require.Equal(t, testCase.Out, out)
		require.Equal(t, "value", ctx.Value(key))
	}
}

func splitString(s string) []string {
	return strings.Split(s, "")
}

func splitStringWithContext(ctx context.Context, s string) []string {
	t := ctx.Value(testKey).(*testing.T)
	require.Equal(t, "value", ctx.Value(key))

	return strings.Split(s, "")
}
