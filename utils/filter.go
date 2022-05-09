package utils

import "context"

// Filter takes an array of T's and applies the callback fn to each element.
// If fn returns true, the element is included in the returned collection, if false it is excluded.
// Example: [1, 2, 3].Filter(<number is odd?>) = [1, 3].
func Filter[T any](arr []T, fn func(T) bool) (out []T) {
	for _, elem := range arr {
		if fn(elem) {
			out = append(out, elem)
		}
	}
	return out
}

// FilterContext is "context aware" and will pass the parent func's ctx param along to the callback fn.
func FilterContext[T any](
	ctx context.Context,
	arr []T,
	fn func(ctx context.Context, elem T) bool,
) (out []T) {
	for _, elem := range arr {
		if fn(ctx, elem) {
			out = append(out, elem)
		}
	}
	return out
}

// FilterError will return early with the first error encountered in the callback fn, if any.
func FilterError[T any](
	arr []T,
	fn func(elem T) (bool, error),
) (out []T, err error) {
	for _, elem := range arr {
		include, err := fn(elem)
		if err != nil {
			return nil, err
		}
		if include {
			out = append(out, elem)
		}
	}
	return out, nil
}

// FilterContextError combines both "FilterContext" and "FilterError" approaches, passing through the ctx,
// and stopping and returning the first error encountered.
func FilterContextError[T any](
	ctx context.Context,
	arr []T,
	fn func(ctx context.Context, elem T) (bool, error),
) (out []T, err error) {
	for _, elem := range arr {
		include, err := fn(ctx, elem)
		if err != nil {
			return nil, err
		}
		if include {
			out = append(out, elem)
		}
	}
	return out, nil
}
