package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindWrongNumber(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	require.Equal(t, 127, FindInvalidNumber(input, 5))
}

func TestFindInvalidTotal(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	require.Equal(t, 62, FindInvalidTotal(input, 127))
}

func TestCheckNumberValid(t *testing.T) {

	require.True(t, CheckNumberValid(40, []int{35, 20, 15, 25, 47}))
	require.False(t, CheckNumberValid(127, []int{95, 102, 117, 150, 182}))
}
