package utils

import (
	"errors"
	"strings"
)

type Result[T any] struct {
	Value T
	Error error
}

type Results[T any] []Result[T]

func (o *Results[T]) Values() (values []T) {
	for _, result := range *o {
		if result.Error == nil {
			values = append(values, result.Value)
		}
	}
	return values
}

func (o *Results[T]) Errors() (errors []error) {
	for _, result := range *o {
		if result.Error != nil {
			errors = append(errors, result.Error)
		}
	}
	return errors
}

func (o *Results[T]) Error() error {
	errs := o.Errors()
	if len(errs) > 0 {
		return errors.New(strings.Join(ErrorStrings(errs), ", "))
	}
	return nil
}
