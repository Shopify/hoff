package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyByField(t *testing.T) {
	type TestExample struct {
		Id   string
		Name string
		Int  int
	}
	first := TestExample{Id: "one", Name: "first", Int: 1}
	second := TestExample{Id: "two", Name: "second", Int: 2}
	third := TestExample{Id: "three", Name: "third", Int: 3}

	items := []TestExample{first, second, third}

	result, err := KeyByField(items, "Id")
	require.NoError(t, err)
	expected := map[string]TestExample{
		"one":   first,
		"two":   second,
		"three": third,
	}
	require.Equal(t, expected, result)

	result, err = KeyByField(items, "Name")
	require.NoError(t, err)
	expected = map[string]TestExample{
		"first":  first,
		"second": second,
		"third":  third,
	}
	require.Equal(t, expected, result)

	result, err = KeyByField(items, "Int")
	require.Error(t, err)
	require.Nil(t, result)
}
