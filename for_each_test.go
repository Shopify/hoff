package hoff

import (
	"context"
	"errors"
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

func TestForEachContextError(t *testing.T) {
	var key = contextKey("key")
	input := []string{"aaa", "bbb"}

	t.Run(
		"success", func(t *testing.T) {
			var calledInputs = make([]string, 0, len(input))
			fn := func(c context.Context, s string) error {
				require.Equal(t, "a_value", c.Value(key))
				calledInputs = append(calledInputs, s)
				return nil
			}
			ctx := context.WithValue(context.Background(), key, "a_value")

			err := ForEachContextError(ctx, input, fn)
			require.Nil(t, err)
			require.Equal(t, input, calledInputs)
		},
	)

	t.Run(
		"failure", func(t *testing.T) {
			var calledInputs = make([]string, 0, len(input))
			fn := func(c context.Context, s string) error {
				require.Equal(t, "a_value", c.Value(key))
				calledInputs = append(calledInputs, s)
				return errors.New("catastrophic error")
			}
			ctx := context.WithValue(context.Background(), key, "a_value")

			err := ForEachContextError(ctx, input, fn)
			require.ErrorContains(t, err, "catastrophic error")
			require.Equal(t, input[0:1], calledInputs) // only first fn was called beacuse it return error
		},
	)

}
