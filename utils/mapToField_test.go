package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapToField(t *testing.T) {
	arr := []StructWithId{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
		{
			Id: "3",
		},
	}

	ids := MapToStringField(arr, "Id")
	require.Equal(t, []string{"1", "2", "3"}, ids)
}

func TestMapToFieldWithOutStructs(t *testing.T) {
	arr := []int{1, 2, 3}

	ids := MapToStringField(arr, "Id")
	require.Equal(t, []string{"", "", ""}, ids)
}
