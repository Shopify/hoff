package hoff

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleResults_HasError() {
	errStrings := Results[int]{
		{0, nil},
		{1, errors.New("first")},
		{2, errors.New("second")},
	}
	fmt.Println(errStrings.HasError())

	noErrors := Results[int]{
		{0, nil},
		{1, nil},
	}
	fmt.Println(noErrors.HasError())
	// Output:
	// true
	// false
}

func ExampleResults_Values() {
	errStrings := Results[int]{
		{0, nil},
		{1, errors.New("first")},
		{2, errors.New("second")},
	}
	fmt.Println(errStrings.Values())

	noErrors := Results[int]{
		{0, nil},
		{1, nil},
	}
	fmt.Println(noErrors.Values())
	// Output:
	// [0]
	// [0 1]
}

func ExampleResults_Errors() {
	errStrings := Results[int]{
		{0, nil},
		{1, errors.New("first")},
		{2, errors.New("second")},
	}
	fmt.Println(errStrings.Errors())

	noErrors := Results[int]{
		{0, nil},
		{1, nil},
	}
	fmt.Println(noErrors.Errors())
	// Output:
	// [first second]
	// []
}

func ExampleResults_Error() {
	errStrings := Results[int]{
		{0, nil},
		{1, errors.New("first")},
		{2, errors.New("second")},
	}
	fmt.Println(errStrings.Error().Error())

	noErrors := Results[int]{
		{0, nil},
		{1, nil},
	}
	fmt.Println(noErrors.Error() == nil, noErrors.Error())
	// Output:
	// first, second
	// true <nil>
}

func TestResults(t *testing.T) {
	ok := Results[int]{
		{1, nil},
		{2, nil},
		{3, nil},
	}
	err := Results[int]{
		{1, nil},
		{2, nil},
		{0, errors.New("fizz")},
	}

	errStrings := Results[int]{
		{0, errors.New("zeroth")},
		{1, errors.New("first")},
		{2, errors.New("second")},
	}

	t.Run(
		"HasErrors", func(t *testing.T) {
			require.False(t, ok.HasError())
			require.True(t, err.HasError())
		},
	)

	t.Run(
		"Errors", func(t *testing.T) {
			require.Equal(t, []error(nil), ok.Errors())
			require.Equal(t, []error{errors.New("fizz")}, err.Errors())
		},
	)

	t.Run(
		"Values", func(t *testing.T) {
			require.Equal(t, []int{1, 2, 3}, ok.Values())
			require.Equal(t, []int{1, 2}, err.Values())
		},
	)

	t.Run(
		"Error", func(t *testing.T) {
			require.Nil(t, ok.Error())
			require.Equal(t, errors.New("fizz"), err.Error())
			require.Equal(t, errors.New("zeroth, first, second"), errStrings.Error())
		},
	)
}
