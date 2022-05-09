package utils

import "context"

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
