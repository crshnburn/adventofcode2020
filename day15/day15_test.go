package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	initial := []int{0, 3, 6}

	game := NewGame(initial)
	require.Equal(t, 3, game.turn)
	require.Equal(t, 6, game.last)
}

func TestTakeTurn(t *testing.T) {
	game := NewGame([]int{0, 3, 6})
	game.TakeTurn()
	require.Equal(t, 4, game.turn)
	require.Equal(t, 0, game.last)
	game.TakeTurn()
	require.Equal(t, 5, game.turn)
	require.Equal(t, 3, game.last)
	game.TakeTurn()
	require.Equal(t, 6, game.turn)
	require.Equal(t, 3, game.last)
	game.TakeTurn()
	require.Equal(t, 7, game.turn)
	require.Equal(t, 1, game.last)
}

func TestGet2020Turn(t *testing.T) {
	game := NewGame([]int{0, 3, 6})
	for game.turn < 2020 {
		game.TakeTurn()
	}
	require.Equal(t, 436, game.last)
}
