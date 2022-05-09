package utils

import (
	"context"
	"fmt"
	"sync"
)

func Map[In, Out any](arr []In, fn func(In) Out) []Out {
	out := make([]Out, 0, len(arr))
	for _, elem := range arr {
		out = append(out, fn(elem))
	}
	return out
}

// MapToValues returns an array of just the values of the input Map.
func MapToValues[In any, Key comparable](arr map[Key]In) []In {
	out := make([]In, 0, len(arr))
	for _, elem := range arr {
		out = append(out, elem)
	}
	return out
}

func MapToSlice[M ~map[K]V, K comparable, V, R any](items M, f func(K, V) R) []R {
	out := make([]R, 0, len(items))
	for k, v := range items {
		out = append(out, f(k, v))
	}
	return out
}

func MapError[In, Out any](arr []In, fn func(In, int) (Out, error)) ([]Out, error) {
	out := make([]Out, 0, len(arr))
	for i, elem := range arr {
		mapped, err := fn(elem, i)
		if err != nil {
			return nil, err
		}
		out = append(out, mapped)
	}

	return out, nil
}

func MapContextError[In, Out any](
	ctx context.Context,
	arr []In,
	fn func(context.Context, In, int) (Out, error),
) ([]Out, error) {
	out := make([]Out, 0, len(arr))
	for i, elem := range arr {
		mapped, err := fn(ctx, elem, i)
		if err != nil {
			return nil, err
		}
		out = append(out, mapped)
	}

	return out, nil
}

type MapConcurrentFunction[T1, T2 any] func(ctx context.Context, t1 T1, i int) (T2, error)

func MapConcurrent[T1, T2 any](ctx context.Context, t1s []T1, fn MapConcurrentFunction[T1, T2]) Results[T2] {
	results := make(Results[T2], len(t1s))
	var wg sync.WaitGroup

	for i, t := range t1s {
		wg.Add(1)
		go func(ctx context.Context, t1 T1, i int) {
			defer wg.Done()
			defer func() {
				r := recover()
				if r != nil {
					var defaultT2 T2
					results[i] = Result[T2]{
						Value: defaultT2,
						Error: fmt.Errorf("MapConcurrent recovered from panic while processing value at index %d: %v", i, r),
					}
				}
			}()

			t2, err := fn(ctx, t1, i)

			// this should be threadsafe since each goroutine only gets
			// one index and we never resize
			results[i] = Result[T2]{
				Value: t2,
				Error: err,
			}
		}(ctx, t, i)
	}
	wg.Wait()

	return results
}
