package main

import "fmt"

type Game struct {
	turn   int
	last   int
	memory map[int][]int
}

func NewGame(initial []int) *Game {
	memory := make(map[int][]int)
	turn := 0
	last := -1
	for _, val := range initial {
		turn++
		memory[val] = []int{turn}
		last = val
	}
	return &Game{turn: turn, last: last, memory: memory}
}

func (g *Game) TakeTurn() {
	g.turn++
	turns, found := g.memory[g.last]
	if !found {
		g.memory[g.last] = []int{g.turn}
		g.last = 0
	} else {
		if len(turns) == 1 {
			g.last = 0
			g.memory[g.last] = append(g.memory[g.last], g.turn)
		} else {
			length := len(turns)
			g.last = turns[length-1] - turns[length-2]
			g.memory[g.last] = append(g.memory[g.last], g.turn)
		}
	}
	// fmt.Println(g.turn, g.last, g.memory)
}

func main() {
	game := NewGame([]int{1, 20, 8, 12, 0, 14})
	for game.turn < 2020 {
		game.TakeTurn()
	}
	fmt.Println("Part 1:", game.last)
	for game.turn < 30000000 {
		game.TakeTurn()
	}
	fmt.Println("Part 2:", game.last)
}
