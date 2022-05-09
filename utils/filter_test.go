package utils

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterInts(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	filtered := Filter(
		ints, func(i int) bool {
			return i%2 == 0
		},
	)
	require.ElementsMatch(t, filtered, []int{2, 4})
}

func TestFilterStrings(t *testing.T) {
	strings := []string{"a", "bb", "ccc", "dddd"}
	filtered := Filter(
		strings, func(s string) bool {
			return len(s) >= 3
		},
	)
	require.ElementsMatch(t, filtered, []string{"ccc", "dddd"})
}

func TestFilterStructs(t *testing.T) {
	type Foo struct {
		Id   int
		Name string
	}
	things := []Foo{
		{Id: 1, Name: "bar"},
		{Id: 2, Name: "baz"},
		{Id: 3, Name: "bat"},
	}

	filtered := Filter(
		things, func(t Foo) bool {
			return t.Id == 2 || t.Name == "bat"
		},
	)
	require.Len(t, filtered, 2)
}

func contextAwareCallbackInt(ctx context.Context, i int) bool {
	// check whether ctx's value for "foo" is "bar" and the number is odd
	return ctx.Value("foo") == "bar" && i%2 == 1
}

func TestFilterContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "foo", "bar")
	ints := []int{1, 2, 3}
	filtered := FilterContext(ctx, ints, contextAwareCallbackInt)
	require.ElementsMatch(t, filtered, []int{1, 3})

	ctx2 := context.Background()
	filtered = FilterContext(ctx2, ints, contextAwareCallbackInt)
	// should not match anything since the context doesn't have a value for "foo"
	require.Len(t, filtered, 0)

	ctx3 := context.WithValue(context.Background(), "foo", "nada")
	filtered = FilterContext(ctx3, ints, contextAwareCallbackInt)
	// should not match anything since the context "foo" value is not "bar"
	require.Len(t, filtered, 0)
}

func contextErrorAwareCallbackInt(ctx context.Context, i int) (bool, error) {
	if ctx.Value("throwError") == true {
		return false, fmt.Errorf("throwError true")
	}
	// check whether ctx's value for "foo" is "bar" and the number is odd
	return ctx.Value("foo") == "bar" && i%2 == 1, nil
}

func TestFilterContextError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "foo", "bar")
	ints := []int{1, 2, 3}
	filtered, err := FilterContextError(ctx, ints, contextErrorAwareCallbackInt)
	require.NoError(t, err)
	require.ElementsMatch(t, filtered, []int{1, 3})

	ctx2 := context.Background()
	filtered, err = FilterContextError(ctx2, ints, contextErrorAwareCallbackInt)
	// should not match anything since the context doesn't have a value for "foo"
	require.NoError(t, err)
	require.Len(t, filtered, 0)

	// create a context with foo = bar AND throwError = true
	ctx3 := context.WithValue(context.Background(), "throwError", true)
	filtered, err = FilterContextError(ctx3, ints, contextErrorAwareCallbackInt)
	// should throw an error
	require.Error(t, err)
	require.Nil(t, filtered)
}
