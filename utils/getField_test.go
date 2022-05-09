package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type GetStringFieldTestCase struct {
	i             interface{}
	prop          string
	expectedValue string
}

type StructWithId struct {
	Id string
}

var getStringFieldTestCases = []GetStringFieldTestCase{
	{
		i: StructWithId{
			Id: "123",
		},
		prop:          "Id",
		expectedValue: "123",
	},
	{
		i: &StructWithId{
			Id: "123",
		},
		prop:          "Id",
		expectedValue: "123",
	},
	{
		i:             234,
		prop:          "Id",
		expectedValue: "",
	},
	{
		i:             nil,
		prop:          "Id",
		expectedValue: "",
	},
	{
		i:             make(chan string),
		prop:          "Id",
		expectedValue: "",
	},
}

func TestGetStringField(t *testing.T) {
	for _, testCase := range getStringFieldTestCases {
		actualValue := GetStringField(testCase.i, testCase.prop)
		require.Equal(t, testCase.expectedValue, actualValue)
	}
}
