package utils

import "context"

// FlatMap applies a transformation to an array of elements and
// returns another array with the transformed result
func FlatMap[In, Out any](arr []In, fn func(In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(elem)...)
	}
	return out
}

// FlatMap applies a transformation to an array of elements and
// returns another array with the transformed result. If one of the
// transformations fails, it will return early
func FlatMapError[In, Out any](arr []In, fn func(In) ([]Out, error)) (out []Out, err error) {
	for _, elem := range arr {
		tr, err := fn(elem)
		if err != nil {
			return nil, err
		}
		out = append(out, tr...)
	}
	return out, nil
}

// FlatMapContext applies the FlatMap transformation while, at the same time,
// shares a context with the transforming function.
func FlatMapContext[In, Out any](ctx context.Context, arr []In, fn func(context.Context, In) []Out) (out []Out) {
	for _, elem := range arr {
		out = append(out, fn(ctx, elem)...)
	}
	return out
}

// FlatMapContext applies the FlatMap transformation while, at the same time,
// shares a context with the transforming function. If one of the
// transformations fails, it will return early
func FlatMapContextError[In, Out any](ctx context.Context, arr []In, fn func(context.Context, In) ([]Out, error)) (out []Out, err error) {
	for _, elem := range arr {
		tr, err := fn(ctx, elem)
		if err != nil {
			return nil, err
		}
		out = append(out, tr...)
	}
	return out, nil
}
