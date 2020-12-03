package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CountTrees(lines []string, right int, down int) int {
	column := 0
	trees := 0
	for i := 0; i < len(lines); i += down {
		if lines[i][column] == '#' {
			trees++
		}
		column = (column + right) % len(lines[i])
	}
	return trees
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1: You hit", CountTrees(lines, 3, 1), "trees")
	part2Trees := CountTrees(lines, 1, 1) *
		CountTrees(lines, 3, 1) *
		CountTrees(lines, 5, 1) *
		CountTrees(lines, 7, 1) *
		CountTrees(lines, 1, 2)
	fmt.Println("Part 2: There are", part2Trees, "trees")
}
