package hoff

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type forEachTestCase struct {
	In  []string
	Out []string
}

var forEachTestCases = []forEachTestCase{
	{
		In:  []string{"aaa", "bbb"},
		Out: []string{"aaa", "bbb"},
	},
	{
		In:  []string{"this", "that"},
		Out: []string{"this", "that"},
	},
}

func TestForEach(t *testing.T) {
	for _, testCase := range forEachTestCases {
		// foreach does not return a value, so we need to test
		// that the receiver function gets called by pushing each value to an array.
		var stringSlice = make([]string, 0, len(testCase.In))
		fn := func(s string) {
			stringSlice = append(stringSlice, s)
		}

		ForEach(testCase.In, fn)
		require.Equal(t, testCase.Out, stringSlice)
	}
}

func TestForEachContext(t *testing.T) {
	type contextKey string
	var key = contextKey("key")

	for _, testCase := range forEachTestCases {
		// foreach does not return a value, so we need to test
		// that the receiver function gets called by pushing each value to an array.
		var stringSlice = make([]string, 0, len(testCase.In))
		fn := func(c context.Context, s string) {
			require.Equal(t, "a_value", c.Value(key))
			stringSlice = append(stringSlice, s)
		}
		ctx := context.WithValue(context.Background(), key, "a_value")

		ForEachContext(ctx, testCase.In, fn)
		require.Equal(t, testCase.Out, stringSlice)
	}
}
