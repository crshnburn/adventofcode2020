package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tile struct {
	q int
	r int
}

func (t *Tile) String() string {
	return fmt.Sprintf("%d,%d", t.q, t.r)
}

func (t *Tile) GetNeighbours() []*Tile {
	return []*Tile{
		&Tile{q: t.q + 1, r: t.r - 1},
		&Tile{q: t.q + 1, r: t.r},
		&Tile{q: t.q, r: t.r + 1},
		&Tile{q: t.q - 1, r: t.r + 1},
		&Tile{q: t.q - 1, r: t.r},
		&Tile{q: t.q, r: t.r - 1},
	}
}

func (t *Tile) Move(instructions string) *Tile {
	instr := instructions
	for len(instr) > 0 {
		switch {
		case instr[0] == 'e':
			instr = instr[1:]
			t.q++
		case instr[0] == 'w':
			instr = instr[1:]
			t.q--
		case instr[0:2] == "se":
			instr = instr[2:]
			t.r++
		case instr[0:2] == "ne":
			instr = instr[2:]
			t.q++
			t.r--
		case instr[0:2] == "nw":
			instr = instr[2:]
			t.r--
		case instr[0:2] == "sw":
			instr = instr[2:]
			t.q--
			t.r++
		}
	}
	return t
}

func CountBlack(instructions []string) (int, map[string]*Tile) {
	blackLocs := make(map[string]*Tile)
	for _, instruction := range instructions {
		tile := &Tile{q: 0, r: 0}
		loc := tile.Move(instruction).String()
		if _, ok := blackLocs[loc]; ok {
			delete(blackLocs, loc)
		} else {
			blackLocs[loc] = tile
		}
	}
	return len(blackLocs), blackLocs
}

func isBlack(blackLoc map[string]*Tile, tile *Tile) bool {
	_, isBlack := blackLoc[tile.String()]
	count := 0
	for _, neighbour := range tile.GetNeighbours() {
		if _, ok := blackLoc[neighbour.String()]; ok {
			count++
		}
	}
	if isBlack {
		return count != 0 && count <= 2
	} else {
		return count == 2
	}
}

func GoL(blackLoc map[string]*Tile) map[string]*Tile {
	newBlackLoc := make(map[string]*Tile)
	for _, tile := range blackLoc {
		if isBlack(blackLoc, tile) {
			newBlackLoc[tile.String()] = tile
		}
		for _, neighbour := range tile.GetNeighbours() {
			if isBlack(blackLoc, neighbour) {
				newBlackLoc[neighbour.String()] = neighbour
			}
		}
	}
	return newBlackLoc
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	count, blackLoc := CountBlack(instructions)
	fmt.Println("Part 1:", count)
	for i := 0; i < 100; i++ {
		blackLoc = GoL(blackLoc)
	}
	fmt.Println("Part 2:", len(blackLoc))
}
