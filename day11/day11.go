package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Seat int

const (
	empty Seat = iota
	occupied
	floor
)

type FloorPlan struct {
	plan          [][]Seat
	width, height int
}

func (f *FloorPlan) String() string {
	var buf bytes.Buffer
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			b := byte(' ')
			switch f.plan[y][x] {
			case occupied:
				b = '#'
			case empty:
				b = 'L'
			case floor:
				b = '.'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func NewFloorPlan(raw []string) *FloorPlan {
	plan := make([][]Seat, len(raw))
	for i, _ := range plan {
		plan[i] = make([]Seat, len(raw[0]))
	}
	for i, row := range raw {
		for j, place := range row {
			switch place {
			case '#':
				plan[i][j] = occupied
			case 'L':
				plan[i][j] = empty
			case '.':
				plan[i][j] = floor
			}
		}
	}
	return &FloorPlan{plan: plan, width: len(raw[0]), height: len(raw)}
}

func (f *FloorPlan) CountSee(x, y int) int {
	count := 0

	for i := x - 1; i >= 0; i-- {
		if f.plan[y][i] == occupied {
			count++
			break
		} else if f.plan[y][i] == empty {
			break
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if f.plan[j][i] == occupied {
			count++
			break
		} else if f.plan[j][i] == empty {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		if f.plan[i][x] == occupied {
			count++
			break
		} else if f.plan[i][x] == empty {
			break
		}
	}

	for i, j := x+1, y-1; i < len(f.plan[0]) && j >= 0; i, j = i+1, j-1 {
		if f.plan[j][i] == occupied {
			count++
			break
		} else if f.plan[j][i] == empty {
			break
		}
	}

	for i := x + 1; i < len(f.plan[0]); i++ {
		if f.plan[y][i] == occupied {
			count++
			break
		} else if f.plan[y][i] == empty {
			break
		}
	}

	for i, j := x+1, y+1; i < len(f.plan[0]) && j < len(f.plan); i, j = i+1, j+1 {
		if f.plan[j][i] == occupied {
			count++
			break
		} else if f.plan[j][i] == empty {
			break
		}
	}

	for i := y + 1; i < len(f.plan); i++ {
		if f.plan[i][x] == occupied {
			count++
			break
		} else if f.plan[i][x] == empty {
			break
		}
	}

	for i, j := x-1, y+1; i >= 0 && j < len(f.plan); i, j = i-1, j+1 {
		if f.plan[j][i] == occupied {
			count++
			break
		} else if f.plan[j][i] == empty {
			break
		}
	}

	return count
}

func (f *FloorPlan) CanSit(x, y int) bool {
	if f.plan[y][x] == empty {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				checkx := x + i
				checky := y + j
				if !(i == 0 && j == 0) && checkx >= 0 && checky >= 0 && checkx < len(f.plan[0]) && checky < len(f.plan) {
					if f.plan[checky][checkx] == occupied {
						return false
					}
				}
			}
		}
		return true
	}
	return false
}

func (f *FloorPlan) NeedVacate(x, y int) bool {
	if f.plan[y][x] == occupied {
		count := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				checkx := x + i
				checky := y + j
				if !(i == 0 && j == 0) && checkx >= 0 && checky >= 0 && checkx < len(f.plan[0]) && checky < len(f.plan) {
					// fmt.Println(checkx, checky, f.plan[checky][checkx])
					if f.plan[checky][checkx] == occupied {
						count++
					}
				}
			}
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

func (f *FloorPlan) Equal(newPlan [][]Seat) bool {
	for i, row := range f.plan {
		for j, val := range row {
			if newPlan[i][j] != val {
				return false
			}
		}
	}
	return true
}

func (f *FloorPlan) Step2() bool {
	plan := make([][]Seat, f.height)
	for i, _ := range plan {
		plan[i] = make([]Seat, f.width)
	}
	for y, row := range f.plan {
		for x, seat := range row {
			switch seat {
			case floor:
				plan[y][x] = floor
			case occupied:
				seen := f.CountSee(x, y)
				// fmt.Println(x, y, seen)
				if seen > 4 {
					plan[y][x] = empty
				} else {
					plan[y][x] = occupied
				}
			case empty:
				seen := f.CountSee(x, y)
				if seen == 0 {
					plan[y][x] = occupied
				} else {
					plan[y][x] = empty
				}
			}
		}
	}
	nochanges := f.Equal(plan)
	f.plan = plan
	return nochanges
}

func (f *FloorPlan) Step() bool {
	plan := make([][]Seat, f.height)
	for i, _ := range plan {
		plan[i] = make([]Seat, f.width)
	}
	for y, row := range f.plan {
		for x, seat := range row {
			switch seat {
			case floor:
				plan[y][x] = floor
			case occupied:
				if f.NeedVacate(x, y) {
					plan[y][x] = empty
				} else {
					plan[y][x] = occupied
				}
			case empty:
				if f.CanSit(x, y) {
					plan[y][x] = occupied
				} else {
					plan[y][x] = empty
				}
			}
		}
	}
	nochanges := f.Equal(plan)
	f.plan = plan
	return nochanges
}

func (f *FloorPlan) CountSeats() int {
	for !f.Step() {

	}
	count := 0
	for _, row := range f.plan {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}
	return count
}

func (f *FloorPlan) CountSeats2() int {
	for !f.Step2() {

	}
	count := 0
	for _, row := range f.plan {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var initial []string
	for scanner.Scan() {
		initial = append(initial, scanner.Text())
	}

	floorPlan := NewFloorPlan(initial)
	fmt.Println("Part 1:", floorPlan.CountSeats())
	floorPlan = NewFloorPlan(initial)
	fmt.Println("Part 2:", floorPlan.CountSeats2())
}
