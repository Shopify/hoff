package utils

import (
	"context"
	"fmt"
)

func sumInts(acc int, elem int, _ int) int {
	return acc + elem
}

func sumIntsError(acc int, elem int, _ int) (int, error) {
	if elem == 4 {
		return acc, fmt.Errorf("invalid element (4) encountered")
	}
	return acc + elem, nil
}

func sumIntsCtxMultiplier(ctx context.Context, acc int, elem int, _ int) int {
	m := ctx.Value("multiplier").(int)
	return acc + (elem * m)
}

func sumIntsCtxDividerError(ctx context.Context, acc int, elem int, _ int) (int, error) {
	m := ctx.Value("divisor").(int)
	if m == 0 {
		return acc, fmt.Errorf("division by zero")
	}
	return acc + (elem / m), nil
}

func ExampleReduce() {
	a := []int{1, 2, 3, 4, 5}
	total := Reduce(a, sumInts, 0)
	fmt.Println(total)
	// Output: 15
}

func ExampleReduceError() {
	a := []int{1, 2, 3}
	fmt.Println(ReduceError(a, sumIntsError, 0))
	// Output: 6 <nil>
}

func ExampleReduceError_error() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(ReduceError(a, sumIntsError, 0))
	// Output: 6 invalid element (4) encountered
}

func ExampleReduceContext() {
	ctx := context.WithValue(context.Background(), "multiplier", 2)
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(ReduceContext(ctx, a, sumIntsCtxMultiplier, 0))
	// Output: 30
}

func ExampleReduceContextError() {
	ctx := context.WithValue(context.Background(), "divisor", 2)
	a := []int{2, 4, 6}
	fmt.Println(ReduceContextError(ctx, a, sumIntsCtxDividerError, 0))
	// Output: 6 <nil>
}

func ExampleReduceContextError_error() {
	ctx := context.WithValue(context.Background(), "divisor", 0)
	a := []int{2, 4, 6}
	fmt.Println(ReduceContextError(ctx, a, sumIntsCtxDividerError, 0))
	// Output: 0 division by zero
}
