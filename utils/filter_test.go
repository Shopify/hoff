package utils

import (
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
