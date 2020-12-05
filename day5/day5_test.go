package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculatePassId(t *testing.T) {
	require.Equal(t, 357, CalculatePassId("FBFBBFFRLR"))
	require.Equal(t, 567, CalculatePassId("BFFFBBFRRR"))
	require.Equal(t, 119, CalculatePassId("FFFBBBFRRR"))
	require.Equal(t, 820, CalculatePassId("BBFFBBFRLL"))
}
