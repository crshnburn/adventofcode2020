package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanSit(t *testing.T) {
	initial := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	floorPlan := NewFloorPlan(initial)
	require.True(t, floorPlan.CanSit(6, 0))
}

func TestNeedVacate(t *testing.T) {
	initial := []string{
		"#.##.##.##",
		"#######.##",
		"#.#.#..#..",
		"####.##.##",
		"#.##.##.##",
		"#.#####.##",
		"..#.#.....",
		"##########",
		"#.######.#",
		"#.#####.##",
	}
	floorPlan := NewFloorPlan(initial)
	require.False(t, floorPlan.NeedVacate(0, 0))
	require.True(t, floorPlan.NeedVacate(1, 1))
	require.True(t, floorPlan.NeedVacate(0, 2))
}

func TestStep(t *testing.T) {
	initial := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	floorPlan := NewFloorPlan(initial)
	require.False(t, floorPlan.Step())
	require.False(t, floorPlan.Step())
	require.False(t, floorPlan.Step())
	require.False(t, floorPlan.Step())
	require.False(t, floorPlan.Step())
	require.True(t, floorPlan.Step())
	fmt.Println(floorPlan.String())
}

func TestStep2(t *testing.T) {
	initial := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	floorPlan := NewFloorPlan(initial)
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.False(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
	require.True(t, floorPlan.Step2())
	fmt.Println(floorPlan.String())
}

func TestCountSeats(t *testing.T) {
	initial := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	floorPlan := NewFloorPlan(initial)
	require.Equal(t, 37, floorPlan.CountSeats())
}
