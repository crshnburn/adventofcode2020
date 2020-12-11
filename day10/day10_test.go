package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoltRange(t *testing.T) {
	adapters := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	noOnes, _, noThrees := JoltRange(adapters)

	require.Equal(t, 7, noOnes)
	require.Equal(t, 5, noThrees)

	adapters = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	noOnes, _, noThrees = JoltRange(adapters)

	require.Equal(t, 22, noOnes)
	require.Equal(t, 10, noThrees)
}

func TestJoltCombo(t *testing.T) {
	adapters := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	// require.Equal(t, 8, JoltCombos(adapters))
	adapters = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	require.Equal(t, 19208, JoltCombos(adapters))
}
