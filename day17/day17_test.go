package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeighbours(t *testing.T) {
	initial := [][]bool{[]bool{false, true, false}, []bool{false, false, true}, []bool{true, true, true}}
	grid := &Grid{layers: [][][]bool{initial}, layerSize: 3, layerDepth: 1}
	require.Equal(t, 1, grid.Neighbours(0, 0, 0))
	require.Equal(t, 1, grid.Neighbours(1, 0, 0))
}

func TestScale(t *testing.T) {
	initial := [][]bool{[]bool{false, true, false}, []bool{false, false, true}, []bool{true, true, true}}
	grid := &Grid{layers: [][][]bool{initial}, layerSize: 3, layerDepth: 1}
	grid.String()
	grid = grid.Scale()
	grid.String()
}

func TestCount(t *testing.T) {
	initial := [][]bool{[]bool{false, true, false}, []bool{false, false, true}, []bool{true, true, true}}
	grid := &Grid{layers: [][][]bool{initial}, layerSize: 3, layerDepth: 1}
	require.Equal(t, 5, grid.Count())

}

func TestCycle(t *testing.T) {
	initial := [][]bool{[]bool{false, true, false}, []bool{false, false, true}, []bool{true, true, true}}
	grid := &Grid{layers: [][][]bool{initial}, layerSize: 3, layerDepth: 1}
	grid = grid.Scale()
	grid = grid.Cycle()
	grid.String()
	require.Equal(t, 11, grid.Count())
	grid = grid.Scale()
	grid = grid.Cycle()
	require.Equal(t, 21, grid.Count())
}
