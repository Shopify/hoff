package utils

import "context"

// FlatMap applies a transformation to an array of elements and
// returns another array with the transformed result
// Example: FlatMap([]string{"abcd", "efg"}, func(s string) []int { return []int{len(s)} }) = [4, 3].
func FlatMap[In, Out any](arr []In, fn func(In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(elem)...)
	}
	return out
}

// FlatMapContext applies the FlatMap transformation while, at the same time,
// shares a context with the transforming function.
func FlatMapContext[In, Out any](ctx context.Context, arr []In, fn func(context.Context, In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(ctx, elem)...)
	}
	return out
}
