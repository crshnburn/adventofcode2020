package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountTrees(t *testing.T) {
	lines := []string{"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#"}

	require.Equal(t, 7, CountTrees(lines, 3, 1))
	require.Equal(t, 2, CountTrees(lines, 1, 1))
	require.Equal(t, 3, CountTrees(lines, 5, 1))
	require.Equal(t, 4, CountTrees(lines, 7, 1))
	require.Equal(t, 2, CountTrees(lines, 1, 2))
}
