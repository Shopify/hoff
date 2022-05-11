package main

import (
	"context"
)

func ForEach[T any](arr []T, fn func(T)) {
	for _, elem := range arr {
		fn(elem)
	}
}

func ForEachContext[T any](ctx context.Context, arr []T, fn func(context.Context, T)) {
	for _, elem := range arr {
		fn(ctx, elem)
	}
}

func ForEachContextError[T any](ctx context.Context, arr []T, fn func(context.Context, T) error) error {
	for _, elem := range arr {
		err := fn(ctx, elem)
		if err != nil {
			return err
		}
	}
	return nil
}
