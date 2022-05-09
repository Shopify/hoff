package maps

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

const benchmarkArrayLength = 10_000

var benchResult []int

func TestToValues(t *testing.T) {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	results := ToValues(nums)
	sort.Ints(results)

	require.Equal(t, results, []int{1, 2, 3})
}

func TestToSlice(t *testing.T) {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fn := func(k string, v int) string {
		return fmt.Sprintf("%d => %s", v, k)
	}
	results := ToSlice(nums, fn)
	sort.Strings(results) // order is not guaranteed for map keys

	require.Equal(t, []string{"1 => one", "2 => two", "3 => three"}, results)
}

func BenchmarkToValues(b *testing.B) {
	arr := make(map[int]int, benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}

	var r []int
	for i := 0; i < b.N; i++ {
		r = ToValues(arr)
	}
	benchResult = r
}

func BenchmarkToSlice(b *testing.B) {
	arr := make(map[int]int, benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	var r []int
	for i := 0; i < b.N; i++ {
		r = ToSlice(arr, func(k, v int) int { return k + v })
	}
	benchResult = r
}

func ExampleToValues() {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	results := ToValues(nums)

	sort.Ints(results) // order is not guaranteed for map keys
	fmt.Println(results)
	// Output: [1 2 3]
}

func ExampleToSlice() {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fn := func(k string, v int) string {
		return fmt.Sprintf("%d in text is %s", v, k)
	}
	results := ToSlice(nums, fn)

	sort.Strings(results) // order is not guaranteed for map keys
	fmt.Println(results)
	// Output: [1 in text is one 2 in text is two 3 in text is three]
}
