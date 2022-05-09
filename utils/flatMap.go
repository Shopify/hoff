package utils

import "context"

func FlatMap[In, Out any](arr []In, fn func(In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(elem)...)
	}
	return out
}

func FlatMapContext[In, Out any](ctx context.Context, arr []In, fn func(context.Context, In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(ctx, elem)...)
	}
	return out
}
