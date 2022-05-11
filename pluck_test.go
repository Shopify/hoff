package hoff

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExamplePluck() {
	fmt.Println(Pluck([]map[string]any{{"foo": 1}, {"bar": 4}}, "bar"))
	// Output: [[<nil>] [4]]
}

func TestPluckEmptyKeys(t *testing.T) {
	output := Pluck(
		[]map[string]any{
			{"foo": 1, "bar": 2, "shop": 3},
			{"foo": 4, "bar": 5, "shop": 6},
		},
	)
	require.Equal(t, [][]any{}, output, "empty result")
}

func TestPluckEmptyMaps(t *testing.T) {
	output := Pluck(
		[]map[string]any{}, "foo", "bar",
	)
	require.Equal(t, [][]any{}, output, "empty result")
}

func TestPluckStringBasic(t *testing.T) {
	output := Pluck(
		[]map[string]any{
			{"foo": 1, "bar": 2, "shop": 3},
			{"foo": 4, "bar": 5, "shop": 6},
		}, "foo", "bar",
	)
	require.Equal(t, [][]any{{1, 2}, {4, 5}}, output, "2 keys extracted")
}

func TestPluckStringNil(t *testing.T) {
	output := Pluck(
		[]map[string]any{
			{"foo": 1, "bar": 2, "shop": 3},
			{"foo": 4, "bar": 5},
		}, "foo", "shop",
	)
	require.Equal(t, [][]any{{1, 3}, {4, nil}}, output, "2 keys extracted with nil value")
}

func TestPluckStringSingleKey(t *testing.T) {
	output := Pluck(
		[]map[string]any{
			{"foo": 1, "bar": 2, "shop": 3},
			{"foo": 4, "bar": 5},
		}, "foo",
	)
	require.Equal(t, [][]any{{1}, {4}}, output, "1 keys extracted")
}
