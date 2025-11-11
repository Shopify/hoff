package hoff

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunkInts(t *testing.T) {
	type TestCase struct {
		Input          []int
		ChunkSize      int
		ExpectedOutput [][]int
		Message        string
	}

	testCases := []TestCase{
		{
			Input:          []int{1, 2, 3, 4, 5},
			ChunkSize:      2,
			ExpectedOutput: [][]int{{1, 2}, {3, 4}, {5}},
			Message:        "5 elements, chunk size 2",
		},
		{
			Input:          []int{5, 2, 4, 1, 3},
			ChunkSize:      2,
			ExpectedOutput: [][]int{{5, 2}, {4, 1}, {3}},
			Message:        "5 elements, chunk size 2, order preserved",
		},
		{
			Input:          []int{1, 2, 3, 4, 5},
			ChunkSize:      5,
			ExpectedOutput: [][]int{{1, 2, 3, 4, 5}},
			Message:        "5 elements, chunk size 5",
		},
		{
			Input:          []int{1, 2, 3, 4, 5},
			ChunkSize:      500,
			ExpectedOutput: [][]int{{1, 2, 3, 4, 5}},
			Message:        "5 elements, chunk size 500",
		},
		{
			Input:          []int{1},
			ChunkSize:      2,
			ExpectedOutput: [][]int{{1}},
			Message:        "1 elements, chunk size 2",
		},
		{
			Input:          []int{1},
			ChunkSize:      0,
			ExpectedOutput: [][]int{},
			Message:        "1 element array, chunk size 0",
		},
		{
			Input:          []int{},
			ChunkSize:      2,
			ExpectedOutput: [][]int{},
			Message:        "empty array, chunk size 2",
		},
	}

	for _, testCase := range testCases {
		output := Chunk(testCase.Input, testCase.ChunkSize)
		require.Equal(t, testCase.ExpectedOutput, output, testCase.Message)
	}
}

func TestChunkStrings(t *testing.T) {
	type TestCase struct {
		Input          []string
		ChunkSize      int
		ExpectedOutput [][]string
		Message        string
	}

	testCases := []TestCase{
		{
			Input:          []string{"a", "b", "c"},
			ChunkSize:      2,
			ExpectedOutput: [][]string{{"a", "b"}, {"c"}},
			Message:        "3 elements, chunk size 2",
		},
	}

	for _, testCase := range testCases {
		output := Chunk(testCase.Input, testCase.ChunkSize)
		require.Equal(t, testCase.ExpectedOutput, output, testCase.Message)
	}
}

func TestChunkStructs(t *testing.T) {
	type ExampleStruct struct {
		Foo string
	}

	type TestCase struct {
		Input          []ExampleStruct
		ChunkSize      int
		ExpectedOutput [][]ExampleStruct
		Message        string
	}

	testCases := []TestCase{
		{
			Input:          []ExampleStruct{{Foo: "a"}, {Foo: "b"}, {Foo: "c"}},
			ChunkSize:      2,
			ExpectedOutput: [][]ExampleStruct{{{Foo: "a"}, {Foo: "b"}}, {{Foo: "c"}}},
			Message:        "3 elements, chunk size 2",
		},
	}

	for _, testCase := range testCases {
		output := Chunk(testCase.Input, testCase.ChunkSize)
		require.Equal(t, testCase.ExpectedOutput, output, testCase.Message)
	}
}

func ExampleChunk() {
	fmt.Println(Chunk([]int{1, 2, 3, 4, 5}, 2))
	// Output: [[1 2] [3 4] [5]]
}

func ExampleChunk_withEmptyArray() {
	fmt.Println(Chunk([]int{}, 2))
	// Output: []
}

func ExampleChunk_zeroBatchSize() {
	fmt.Println(Chunk([]int{1, 2, 3}, 0))
	// Output: []
}
