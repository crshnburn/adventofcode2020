package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func CalculatePassId(pass string) int {
	lowRow := 0
	highRow := 127
	lowCol := 0
	highCol := 7
	for _, char := range pass {
		switch char {
		case 'F':
			highRow = highRow - ((highRow - lowRow) / 2) - 1
		case 'B':
			lowRow = lowRow + ((highRow - lowRow) / 2) + 1
		case 'R':
			lowCol = lowCol + ((highCol - lowCol) / 2) + 1
		case 'L':
			highCol = highCol - ((highCol - lowCol) / 2) - 1
		}
	}
	return lowRow*8 + lowCol
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	highest := 0
	seatIds := []int{}
	for scanner.Scan() {
		id := CalculatePassId(scanner.Text())
		seatIds = append(seatIds, id)
		if id > highest {
			highest = id
		}
	}
	fmt.Println("The highest ID is", highest)

	sort.Ints(seatIds)
	for i, seat := range seatIds {
		if seat > 7 {
			if seatIds[i+1] == seat+2 {
				fmt.Println("My seat is", seat+1)
				break
			}
		}
	}
}
