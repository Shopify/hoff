package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// Map implements the basic map function.
func Map[In, Out any](arr []In, fn func(In) Out) []Out {
	out := make([]Out, 0, len(arr))
	for _, elem := range arr {
		out = append(out, fn(elem))
	}
	return out
}

// MapError is the same as `Map` but for functions that might return an error.
func MapError[In, Out any](arr []In, fn func(In) (Out, error)) ([]Out, error) {
	out := make([]Out, 0, len(arr))
	for i, elem := range arr {
		mapped, err := fn(elem)
		if err != nil {
			return nil, fmt.Errorf("MapError got an error in index %d, value %v: %w", i, elem, err)
		}
		out = append(out, mapped)
	}
	return out, nil
}

// MapContextError is the same as `MapError` but with a context.
func MapContextError[In, Out any](
	ctx context.Context,
	arr []In,
	fn func(context.Context, In) (Out, error),
) ([]Out, error) {
	out := make([]Out, 0, len(arr))
	for i, elem := range arr {
		mapped, err := fn(ctx, elem)
		if err != nil {
			return nil, fmt.Errorf("MapContextError got an error in index %d, value %v: %w", i, elem, err)
		}
		out = append(out, mapped)
	}
	return out, nil
}

// MapConcurrent is the same as `MapContextError` but applied with concurrency.
// In spite of concurency, order is guaranteed.
func MapConcurrent[In, Out any](
	ctx context.Context,
	arr []In,
	fn func(ctx context.Context, elem In) (Out, error),
) Results[Out] {
	results := make(Results[Out], len(arr))
	var wg sync.WaitGroup
	for i, t := range arr {
		wg.Add(1)
		go func(ctx context.Context, elem In, i int) {
			defer wg.Done()
			defer func() {
				r := recover()
				if r != nil {
					var val Out
					results[i] = Result[Out]{
						Value: val,
						Error: fmt.Errorf("MapConcurrent recovered from panic while processing index %d, value %v: %v", i, elem, r),
					}
				}
			}()

			val, err := fn(ctx, elem)
			if err != nil {
				err = fmt.Errorf("MapConcurrent got an error in index %d, value %v: %w", i, elem, err)
			}
			results[i] = Result[Out]{val, err}
		}(ctx, t, i)
	}
	wg.Wait()

	return results
}

// MapConcurrentError is the same as `MapConcurrentError` but returns only the
// values and, if it happens, the first error.
func MapConcurrentError[In, Out any](
	ctx context.Context,
	arr []In,
	fn func(ctx context.Context, elem In) (Out, error),
) ([]Out, error) {
	results := make([]Out, len(arr))
	errs := make(chan error)
	defer close(errs)

	var shutdown uint32 // used to gracefuly shutdown in case of error or panic
	for i, elem := range arr {
		go func(ctx context.Context, elem In, i int) {
			defer func() {
				r := recover()
				if r != nil && atomic.CompareAndSwapUint32(&shutdown, 0, 1) {
					errs <- fmt.Errorf("MapConcurrentError recovered from panic while processing index %d, value %v: %v", i, elem, r)
				}
			}()

			r, err := fn(ctx, elem)
			if err != nil && atomic.CompareAndSwapUint32(&shutdown, 0, 1) {
				errs <- fmt.Errorf("MapConcurrentError got an error in index %d, value %v: %w", i, elem, err)
			}
			results[i] = r
			if atomic.LoadUint32(&shutdown) == 0 {
				errs <- nil
			}
		}(ctx, elem, i)
	}

	for range arr {
		err := <-errs
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
