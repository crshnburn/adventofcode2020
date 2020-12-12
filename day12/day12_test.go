package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInstruction(t *testing.T) {
	instr := "N10"
	dir, amount := ParseInstruction(instr)
	require.Equal(t, 'N', dir)
	require.Equal(t, 10, amount)
}

func TestCalculatePos(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	ship := &Ship{north: 0, east: 0, dir: 'E'}
	ship.FollowCourse(instructions)
	require.Equal(t, 25, Abs(ship.north)+Abs(ship.east))
}

func TestCalculatePos2(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	waypoint := &WayPoint{north: 1, east: 10}
	ship := &Ship2{north: 0, east: 0, waypoint: waypoint}
	ship.FollowInstructions(instructions)
	require.Equal(t, 286, Abs(ship.north)+Abs(ship.east))
}
