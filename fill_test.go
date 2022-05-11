package hoff

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type FillTestCase struct {
	Num int
}

var fillTestCases = []FillTestCase{
	{
		Num: 1,
	},
	{
		Num: 666,
	},
	{
		Num: 10000,
	},
}

func TestFill(t *testing.T) {
	for _, fillTestCases := range fillTestCases {
		arr := Fill(1, fillTestCases.Num)
		require.Equal(t, fillTestCases.Num, len(arr))
		for _, elem := range arr {
			require.Equal(t, 1, elem)
		}
	}
}
