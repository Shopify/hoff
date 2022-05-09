package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type contextKey string

const (
	ctxKey               = contextKey("answer")
	benchmarkArrayLength = 10_000
)

var (
	benchErr    error
	benchResult []int
	ctx         = context.WithValue(context.Background(), ctxKey, 42)
)

func TestMap(t *testing.T) {
	results := Map([]string{"a", "b", "c"}, strings.ToUpper)
	require.Equal(t, []string{"A", "B", "C"}, results)
}

func TestMapError(t *testing.T) {
	fn := func(n int) (int, error) {
		if n > 3 {
			return 0, fmt.Errorf("%d is greater than 3", n)
		}
		return n, nil
	}

	t.Run("success", func(t *testing.T) {
		results, err := MapError([]int{1, 2, 3}, fn)
		require.NoError(t, err)
		require.Equal(t, []int{1, 2, 3}, results)
	})

	t.Run("failure", func(t *testing.T) {
		results, err := MapError([]int{2, 3, 4}, fn)
		require.Error(t, err)
		require.Nil(t, results)
	})
}

func TestMapContextError(t *testing.T) {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	t.Run("success", func(t *testing.T) {
		results, err := MapContextError(ctx, []int{1, 2, 3}, fn)
		require.NoError(t, err)
		require.Equal(t, results, []int{1, 2, 3})
	})

	t.Run("failure", func(t *testing.T) {
		results, err := MapContextError(ctx, []int{1, 2, 42}, fn)
		require.Error(t, err)
		require.Nil(t, results)
	})
}

func TestMapConcurrent(t *testing.T) {
	err := errors.New("I can't even")

	t.Run("success", func(t *testing.T) {
		fn := func(ctx context.Context, val int) (string, error) {
			return strconv.Itoa(val), nil
		}
		results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

		expected := Results[string]{
			{"1", nil},
			{"2", nil},
			{"3", nil},
		}
		require.Equal(t, expected, results)
	})

	t.Run("failure", func(t *testing.T) {
		fn := func(ctx context.Context, val int) (string, error) {
			if val == 2 {
				return "", err
			}
			return strconv.Itoa(val), nil
		}

		results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

		expected := Results[string]{
			{"1", nil},
			{"", fmt.Errorf("MapConcurrent got an error in index 1, value 2: %w", err)},
			{"3", nil},
		}
		require.Equal(t, expected, results)
	})

	t.Run("panic", func(t *testing.T) {
		fn := func(ctx context.Context, val int) (string, error) {
			if val == 2 {
				return "", err
			}
			return strconv.Itoa(val), nil
		}

		results := MapConcurrent(context.Background(), []int{1, 2, 3}, fn)

		expected := Results[string]{
			{"1", nil},
			{"", fmt.Errorf("MapConcurrent got an error in index 1, value 2: %w", err)},
			{"3", nil},
		}
		require.Equal(t, expected, results)
	})
}

func TestMapConcurrentError(t *testing.T) {
	err := errors.New("I can't even")

	t.Run("success", func(t *testing.T) {
		fn := func(ctx context.Context, val int) (string, error) {
			return strconv.Itoa(val), nil
		}
		results, err := MapConcurrentError(context.Background(), []int{1, 2, 3}, fn)

		require.Equal(t, []string{"1", "2", "3"}, results)
		require.Nil(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		fn := func(ctx context.Context, val int) (string, error) {
			if val == 2 {
				return "", err
			}
			return strconv.Itoa(val), nil
		}

		results, err := MapConcurrentError(context.Background(), []int{1, 2, 3}, fn)

		require.Error(t, err)
		require.Nil(t, results)
	})

	t.Run("panic", func(t *testing.T) {
		err := errors.New("I can't even")
		fn := func(ctx context.Context, val int) (string, error) {
			if val == 2 {
				return "", err
			}
			return strconv.Itoa(val), nil
		}

		results, err := MapConcurrentError(context.Background(), []int{1, 2, 3}, fn)

		require.Error(t, err)
		require.Nil(t, results)
	})
}

func BenchmarkMap(b *testing.B) {
	arr := make([]int, benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	fn := func(n int) int { return n ^ 2 }

	var r []int
	for i := 0; i < b.N; i++ {
		r = Map(arr, fn)
	}
	benchResult = r
}

func BenchmarkMapError(b *testing.B) {
	arr := make([]int, benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	fn := func(n int) (int, error) {
		if n%2 != 0 {
			return 0, fmt.Errorf("i can't even")
		}
		return n, nil
	}

	var r []int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = MapError(arr, fn)
	}
	benchResult = r
	benchErr = err
}

func BenchmarkMapContextError(b *testing.B) {
	arr := make([]int, benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	fn := func(ctx context.Context, n int) (int, error) {
		if n == ctx.Value(ctxKey).(int) {
			return 0, fmt.Errorf("%d is not the answer", n)
		}
		return n, nil
	}

	var r []int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = MapContextError(ctx, arr, fn)
	}
	benchResult = r
	benchErr = err
}

func BenchmarkMapConcurrent(b *testing.B) {
	arr := make([]int, 2*benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	arr[b.N/2] = 42
	fn := func(_ context.Context, n int) (int, error) {
		if n == 42 {
			return 0, fmt.Errorf("%d is the answer", n)
		}
		return n ^ 2, nil
	}

	var r Results[int]
	for i := 0; i < b.N; i++ {
		r = MapConcurrent(ctx, arr, fn)
	}
	benchResult = r.Values()
	benchErr = r.Errors()[0]
}

func BenchmarkMapConcurrentError(b *testing.B) {
	arr := make([]int, 2*benchmarkArrayLength)
	for i := range arr {
		arr[i] = rand.Int()
	}
	arr[b.N/2] = 42
	fn := func(_ context.Context, n int) (int, error) {
		if n == 42 {
			return 0, fmt.Errorf("%d is the answer", n)
		}
		return n ^ 2, nil
	}

	var r []int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = MapConcurrentError(ctx, arr, fn)
	}
	benchResult = r
	benchErr = err
}

func ExampleMap() {
	arr := []int{0, 1, 2, 3}
	fn := func(n int) int { return n * 2 }
	results := Map(arr, fn)
	fmt.Println(results)
	// Output: [0 2 4 6]
}

func ExampleMapError_success() {
	fn := func(n int) (int, error) {
		if n > 3 {
			return 0, fmt.Errorf("%d is greater than 3", n)
		}
		return n, nil
	}

	results, _ := MapError([]int{1, 2, 3}, fn)

	fmt.Println(results)
	// Output: [1 2 3]
}

func ExampleMapError_failure() {
	fn := func(n int) (int, error) {
		if n > 3 {
			return 0, fmt.Errorf("%d is greater than 3", n)
		}
		return n, nil
	}

	_, err := MapError([]int{2, 3, 4}, fn)

	fmt.Println(err.Error())
	// Output: MapError got an error in index 2, value 4: 4 is greater than 3
}

func ExampleMapContextError_success() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	results, _ := MapContextError(ctx, []int{1, 2, 3}, fn)

	fmt.Println(results)
	// Output: [1 2 3]
}

func ExampleMapContextError_failure() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	_, err := MapContextError(ctx, []int{1, 2, 42}, fn)

	fmt.Println(err.Error())
	// Output: MapContextError got an error in index 2, value 42: 42 is already the answer
}

func ExampleMapConcurrent_success() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	results := MapConcurrent(ctx, []int{1, 2, 3}, fn)

	fmt.Println(results.Values())
	// Output: [1 2 3]
}

func ExampleMapConcurrent_failure() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	results := MapConcurrent(ctx, []int{1, 2, 42}, fn)

	fmt.Println(results.Errors())
	// Output: [<nil> <nil> MapConcurrent got an error in index 2, value 42: 42 is already the answer]
}

func ExampleMapConcurrent_panic() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			panic(fmt.Sprintf("%d is already the answer", n))
		}
		return n, nil
	}

	results := MapConcurrent(ctx, []int{1, 2, 42}, fn)

	fmt.Println(results.Errors())
	// Output: [<nil> <nil> MapConcurrent recovered from panic while processing index 2, value 42: 42 is already the answer]
}

func ExampleMapConcurrentError_success() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	results, _ := MapConcurrentError(ctx, []int{1, 2, 3}, fn)

	fmt.Println(results)
	// Output: [1 2 3]
}

func ExampleMapConcurrentError_failure() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			return 0, fmt.Errorf("%d is already the answer", n)
		}
		return n, nil
	}

	_, err := MapConcurrentError(ctx, []int{1, 2, 42}, fn)

	fmt.Println(err.Error())
	// Output: MapConcurrentError got an error in index 2, value 42: 42 is already the answer
}

func ExampleMapConcurrentError_panic() {
	fn := func(ctx context.Context, n int) (int, error) {
		diff := ctx.Value(ctxKey).(int) - n
		if diff == 0 {
			panic(fmt.Sprintf("%d is already the answer", n))
		}
		return n, nil
	}

	_, err := MapConcurrentError(ctx, []int{1, 2, 42}, fn)

	fmt.Println(err.Error())
	// Output: MapConcurrentError recovered from panic while processing index 2, value 42: 42 is already the answer
}
