package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"

	require.Equal(t, 71, Calculate(input))
}

func TestCalculate2(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"

	require.Equal(t, 231, Calculate2(input))
}
