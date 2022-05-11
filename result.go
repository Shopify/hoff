package hoff

// Result holds a result of a concurrent operation.
type Result[T any] struct {
	Value T
	Error error
}

// Results hold a sequence concurrent operation results in order.
type Results[T any] []Result[T]

// HasError returns whether the sequence of results has any errors.
func (rs Results[T]) HasError() bool {
	for _, r := range rs {
		if r.Error != nil {
			return true
		}
	}
	return false
}

// Values returns an array with the values.
func (rs Results[T]) Values() []T { return Map(rs, func(r Result[T]) T { return r.Value }) }

// Errors returns an array with the errors.
func (rs Results[T]) Errors() []error { return Map(rs, func(r Result[T]) error { return r.Error }) }
