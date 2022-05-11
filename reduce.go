package main

import "context"

// Reduce takes an array of input items, runs the callback on each one and accumulates the result.
func Reduce[T, Acc any](
	arr []T,
	fn func(acc Acc, elem T, index int) Acc,
	acc Acc,
) Acc {
	for i, elem := range arr {
		acc = fn(acc, elem, i)
	}
	return acc
}

// ReduceContext passes the context arg through to the reducer fn.
func ReduceContext[T, Acc any](
	ctx context.Context,
	arr []T,
	fn func(ctx context.Context, acc Acc, elem T, index int) Acc,
	acc Acc,
) Acc {
	for i, elem := range arr {
		acc = fn(ctx, acc, elem, i)
	}
	return acc
}

// ReduceError will stop the reducer when an error is encountered and return the Acc and the error encountered.
func ReduceError[T, Acc any](
	arr []T,
	fn func(acc Acc, elem T, index int) (Acc, error),
	acc Acc,
) (Acc, error) {
	var err error
	for i, elem := range arr {
		acc, err = fn(acc, elem, i)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// ReduceContextError combines both ReduceContext and ReduceError.
func ReduceContextError[T, Acc any](
	ctx context.Context,
	arr []T,
	fn func(ctx context.Context, acc Acc, elem T, index int) (Acc, error),
	acc Acc,
) (Acc, error) {
	var err error
	for i, elem := range arr {
		acc, err = fn(ctx, acc, elem, i)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}
