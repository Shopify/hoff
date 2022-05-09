package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type FlatMapTestCase struct {
	In  []string
	Out []string
	Fn  func(s string) []string
}

var testCases = []FlatMapTestCase{
	{
		In:  []string{"abcd"},
		Out: []string{"a", "b", "c", "d"},
		Fn:  splitString,
	},
	{
		In:  []string{"abcd", "efg"},
		Out: []string{"a", "b", "c", "d", "e", "f", "g"},
		Fn:  splitString,
	},
}

func TestFlatMap(t *testing.T) {
	for _, testCase := range testCases {
		out := FlatMap(testCase.In, testCase.Fn)
		require.Equal(t, testCase.Out, out)
	}
}

func splitString(s string) []string {
	return strings.Split(s, "")
}
