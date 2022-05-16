package hoff

import (
	"errors"
	"strings"
)

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

// Values returns an array with the values of all non-error Results.
func (rs Results[T]) Values() (values []T) {
	for _, result := range rs {
		if result.Error == nil {
			values = append(values, result.Value)
		}
	}
	return values
}

// Errors returns an array with the errors of all non-valid Results.
func (rs Results[T]) Errors() (errors []error) {
	for _, result := range rs {
		if result.Error != nil {
			errors = append(errors, result.Error)
		}
	}
	return errors
}

// Error returns nil if the results contain no errors,
// or an error with the message a comma-separated string of all the errors in the collection.
func (rs Results[T]) Error() error {
	errs := rs.Errors()
	if len(errs) > 0 {
		errorStrings := Map(
			errs, func(err error) string {
				return err.Error()
			},
		)

		return errors.New(strings.Join(errorStrings, ", "))
	}
	return nil
}
