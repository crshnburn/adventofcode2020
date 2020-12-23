package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNextThree(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	require.Equal(t, []int{8, 9, 1}, GetNextThree(circle, 3))
	require.Equal(t, []int{3, 8, 9}, GetNextThree(circle, 7))
}

func TestFindPos(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	for i, cup := range circle {
		require.Equal(t, i, findPos(circle, cup))
	}
}

func TestRemoveFound(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	require.Equal(t, []int{3, 2, 5, 4, 6, 7}, RemoveFound(circle, 3))
}

func TestRemoveCup(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	require.Equal(t, []int{3, 8, 9, 1, 2, 5, 4, 6}, removeCup(circle, 7))
}

func TestFindInsert(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	cup := 3

	circle = RemoveFound(circle, cup)

	require.Equal(t, 1, FindInsert(circle, cup))
}

func TestInsertCups(t *testing.T) {
	circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	cup := 3

	removed := GetNextThree(circle, cup)
	circle = RemoveFound(circle, cup)

	require.Equal(t, []int{3, 2, 8, 9, 1, 5, 4, 6, 7}, InsertCups(circle, cup, removed))
}

func TestGame(t *testing.T) {
	//circle := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	require.Equal(t, "92658374", answerString(Game([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10)))
	require.Equal(t, "67384529", answerString(Game([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 100)))
}

func TestRingGame(t *testing.T) {
	require.Equal(t, "92658374", part1Answer(RingGame([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10)))
	require.Equal(t, "67384529", part1Answer(RingGame([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 100)))
}
