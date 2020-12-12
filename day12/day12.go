package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Ship struct {
	north int
	east  int
	dir   rune
}

type Ship2 struct {
	north    int
	east     int
	waypoint *WayPoint
}

type WayPoint struct {
	north int
	east  int
}

func (s *Ship) Move(dir rune, amount int) {
	switch dir {
	case 'N':
		s.north += amount
	case 'S':
		s.north -= amount
	case 'E':
		s.east += amount
	case 'W':
		s.east -= amount
	}
}

func (s *Ship) Turn(dir rune, amount int) {
	if dir == 'L' {
		for i := 90; i <= amount; i += 90 {
			switch s.dir {
			case 'N':
				s.dir = 'W'
			case 'W':
				s.dir = 'S'
			case 'S':
				s.dir = 'E'
			case 'E':
				s.dir = 'N'
			}
		}
	} else {
		for i := 90; i <= amount; i += 90 {
			switch s.dir {
			case 'N':
				s.dir = 'E'
			case 'W':
				s.dir = 'N'
			case 'S':
				s.dir = 'W'
			case 'E':
				s.dir = 'S'
			}
		}
	}
}

func (w *WayPoint) rotate(dir rune, amount int) {
	if dir == 'R' {
		for i := 90; i <= amount; i += 90 {
			w.north, w.east = -w.east, w.north
		}
	} else {
		for i := 90; i <= amount; i += 90 {
			w.north, w.east = w.east, -w.north
		}
	}
}

func (w *WayPoint) move(dir rune, amount int) {
	switch dir {
	case 'N':
		w.north += amount
	case 'E':
		w.east += amount
	case 'S':
		w.north -= amount
	case 'W':
		w.east -= amount
	}

}

func (s *Ship) Forward(amount int) {
	s.Move(s.dir, amount)
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (s *Ship) FollowCourse(instructions []string) {
	for _, instr := range instructions {
		dir, amount := ParseInstruction(instr)
		switch dir {
		case 'N', 'W', 'E', 'S':
			s.Move(dir, amount)
		case 'F':
			s.Forward(amount)
		case 'L', 'R':
			s.Turn(dir, amount)
		}
	}
}

func (s *Ship2) FollowInstructions(instructions []string) {
	for _, instr := range instructions {
		dir, amount := ParseInstruction(instr)
		switch dir {
		case 'N', 'W', 'E', 'S':
			s.waypoint.move(dir, amount)
		case 'R', 'L':
			s.waypoint.rotate(dir, amount)
		case 'F':
			s.north += amount * s.waypoint.north
			s.east += amount * s.waypoint.east
		}
	}
}

func ParseInstruction(instr string) (rune, int) {
	var dir rune
	var amount int
	fmt.Sscanf(instr, "%c%d", &dir, &amount)
	return dir, amount
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	ship := &Ship{north: 0, east: 0, dir: 'E'}
	ship.FollowCourse(instructions)
	fmt.Println("Part 1:", Abs(ship.north)+Abs(ship.east))
	waypoint := &WayPoint{north: 1, east: 10}
	ship2 := &Ship2{north: 0, east: 0, waypoint: waypoint}
	ship2.FollowInstructions(instructions)
	fmt.Println("Part 2:", Abs(ship2.north)+Abs(ship2.east))
}
